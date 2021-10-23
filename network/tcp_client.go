package network

import "net"

type TcpClient struct {
	dialer *net.Dialer
}

func NewTcpClient() *TcpClient {
	return &TcpClient{}
}

func (cli *TcpClient) Dial(addr string, opts ...interface{}) (*TcpConn, error) {
	var conn net.Conn
	var err error

	if cli.dialer != nil {
		conn, err = cli.dialer.Dial("tcp", addr)
	} else {
		conn, err = net.Dial("tcp", addr)
	}
	if err != nil {
		return nil, err
	}

	var connClose ConnCloseCB
	if len(opts) > 0 {
		connClose = opts[0].(ConnCloseCB)
	}
	var onMsg OnMsgCB
	if len(opts) > 1 {
		onMsg = opts[1].(OnMsgCB)
	}
	bak, err := NewTcpConn(conn, connClose, onMsg)
	if err == nil {
		bak.Run()
	}

	return bak, err
}
