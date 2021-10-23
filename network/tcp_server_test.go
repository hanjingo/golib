package network

import (
	"fmt"
	"testing"
	"time"
)

var addr = "127.0.0.1:10086"

func cli() {
	fmt.Printf("client start\n")
	var close_conn ConnCloseCB = func(c SessionI) {
		fmt.Printf("cli conn close\n")
	}
	var onmsg OnMsgCB = func(c SessionI, data []byte) {
		fmt.Printf("cli recv:%s\n", string(data))
	}

	cli := NewTcpClient()
	c, err := cli.Dial(addr, close_conn, onmsg)
	if err != nil {
		fmt.Printf("create conn fail with err:%v\n", err)
		return
	}
	time.Sleep(time.Duration(10) * time.Millisecond)
	c.Write([]byte("hello work"))
	time.Sleep(time.Duration(3) * time.Second)
}

func serv() {
	fmt.Printf("server start\n")
	var new_conn NewConnCB = func(c SessionI) {
		fmt.Printf("new serv conn\n")
		go c.Run()
	}
	var close_conn ConnCloseCB = func(c SessionI) {
		fmt.Printf("serv conn close\n")
	}
	var onmsg OnMsgCB = func(c SessionI, data []byte) {
		fmt.Printf("serv recv:%s\n", string(data))
	}

	srv := NewTcpServer()
	srv.Listen(addr, new_conn, close_conn, onmsg)
}

// go test -v tcp_client.go tcp_server.go tcp_conn.go tcp_server_test.go define.go interface.go -test.run TestNewTcpServer
func TestNewTcpServer(t *testing.T) {
	go serv()
	time.Sleep(time.Duration(1) * time.Second)
	go cli()
	time.Sleep(time.Duration(2) * time.Second)
}
