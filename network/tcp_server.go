package network

import (
	"net"
)

type TcpServer struct {
	addr string
	li   net.Listener
}

func NewTcpServer() *TcpServer {
	return &TcpServer{}
}

func (s *TcpServer) Listen(addr string, opts ...interface{}) error {
	var err error
	var newConn NewConnCB
	if len(opts) > 0 {
		newConn = opts[0].(NewConnCB)
	}
	var connClose ConnCloseCB
	if len(opts) > 1 {
		connClose = opts[1].(ConnCloseCB)
	}
	var onMsg OnMsgCB
	if len(opts) > 2 {
		onMsg = opts[2].(OnMsgCB)
	}

	if s.li, err = net.Listen("tcp", s.addr); s.li == nil || err != nil {
		return err
	}
	for {
		c, err := s.li.Accept()
		if err != nil {
			return err
		}
		conn, err := NewTcpConn(c, connClose, onMsg)
		if err != nil {
			continue
		}
		go newConn(conn)
	}
}

func (s *TcpServer) Close() {
	if s.li != nil {
		s.li.Close()
	}
}
