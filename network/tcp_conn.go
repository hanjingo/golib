package network

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"
	"runtime"
	"sync"
	"sync/atomic"
)

type TcpConn struct {
	mu   *sync.RWMutex
	conn net.Conn
	stat int32

	readBuf  *bytes.Buffer
	writeBuf *bytes.Buffer

	cancel  context.CancelFunc
	closeCb ConnCloseCB
	onMsgCb OnMsgCB
}

func NewTcpConn(conn net.Conn, opts ...interface{}) (*TcpConn, error) {
	if conn == nil {
		return nil, errors.New("conn cannot be empty")
	}
	var onclose ConnCloseCB
	if len(opts) > 0 {
		onclose = opts[0].(ConnCloseCB)
	}
	var onmsg OnMsgCB
	if len(opts) > 1 {
		onmsg = opts[1].(OnMsgCB)
	}

	back := &TcpConn{
		mu:   new(sync.RWMutex),
		conn: conn,
		stat: INITED,

		readBuf:  new(bytes.Buffer),
		writeBuf: new(bytes.Buffer),

		closeCb: onclose,
		onMsgCb: onmsg,
	}

	// disenable negal algorithm
	back.conn.(*net.TCPConn).SetNoDelay(true)
	return back, nil
}

func (c *TcpConn) Run() error {
	if !atomic.CompareAndSwapInt32(&c.stat, INITED, CONNECTED) {
		return errors.New("already started")
	}

	// catch panic
	defer func() {
		if err := recover(); err != nil {
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			fmt.Printf("%v\n", string(buf[:n]))
		}
	}()
	ctx, cancel := context.WithCancel(context.Background())
	c.cancel = cancel

	// start read corountine
	go c.goRead(ctx)

	// start write corountine
	go c.goWrite(ctx)
	return nil
}

func (c *TcpConn) Destroy() {
	if atomic.CompareAndSwapInt32(&c.stat, CONNECTED, DESTROYED) {
		c.Close()
	}
	// reset buf
	c.writeBuf.Reset()
	c.readBuf.Reset()
}

func (c *TcpConn) Read(arg []byte) (int, error) {
	if atomic.LoadInt32(&c.stat) != CONNECTED &&
		atomic.LoadInt32(&c.stat) != WRITE_CLOSED {
		return 0, errors.New("connection invalid")
	}

	return c.readBuf.Read(arg)
}

func (c *TcpConn) Write(args ...[]byte) (int, error) {
	if atomic.LoadInt32(&c.stat) != CONNECTED {
		return 0, errors.New("connection invalid")
	}

	l := 0
	for _, arg := range args {
		n, err := c.writeBuf.Write(arg)
		if err != nil {
			return l, err
		}
		l += n
	}
	return l, nil
}

func (c *TcpConn) Close() {
	if !atomic.CompareAndSwapInt32(&c.stat, CONNECTED, WRITE_CLOSED) {
		return
	}
	// close connection
	if c.conn != nil {
		c.conn.Close()
	}
	// exit co
	if c.cancel != nil {
		c.cancel()
	}
	if c.closeCb != nil {
		c.closeCb(c)
	}
}

func (c *TcpConn) goRead(ctx context.Context) {
	capa := 1
	if ReadBufSize > 0 {
		capa = ReadBufSize / PackSize
	}
	readC := make(chan []byte, capa)
	freeC := make(chan []byte, capa)
	for i := 0; i < capa; i++ {
		freeC <- make([]byte, PackSize)
	}
	defer func() {
		for len(readC) > 0 {
			c.readBuf.Write(<-readC)
		}
		close(readC)
		close(freeC)
		c.Close()
	}()
	// do loop
	for {
		select {
		case <-ctx.Done():
			return
		case arg := <-readC:
			if c.onMsgCb != nil {
				c.onMsgCb(c, arg)
			} else {
				c.readBuf.Write(arg)
			}
			freeC <- arg
		case tmp := <-freeC:
			n, err := c.conn.Read(tmp) // block
			if err != nil {
				return
			}
			readC <- tmp[:n]
		}
	}
}

func (c *TcpConn) goWrite(ctx context.Context) {
	capa := 1
	if WriteBufSize > 0 {
		capa = WriteBufSize / PackSize
	}
	writeC := make(chan []byte, capa)
	freeC := make(chan []byte, capa)
	for i := 0; i < capa; i++ {
		freeC <- make([]byte, capa)
	}
	defer func() {
		close(writeC)
		close(freeC)
		c.Close()
	}()
	for {
		select {
		case <-ctx.Done():
			for c.writeBuf.Len() > 0 {
				for len(writeC) > 0 {
					if _, err := c.conn.Write(<-writeC); err != nil {
						return
					}
				}
				writeC <- c.writeBuf.Next(PackSize)
			}
			return
		case arg := <-writeC:
			if _, err := c.conn.Write(arg); err != nil {
				return
			}
			freeC <- arg
		case arg := <-freeC:
			n, err := c.writeBuf.Read(arg)
			if err != nil {
				return
			}
			writeC <- arg[:n]
		}
	}
}
