package network

import (
	"time"
)

type ConnCloseCB func(SessionI)
type NewConnCB func(SessionI)
type OnMsgCB func(SessionI, []byte)

// connection int32us
const (
	UNKNOWN      int32 = 0
	INITED       int32 = 1
	CONNECTED    int32 = 2
	WRITE_CLOSED int32 = 3
	READ_CLOSED  int32 = 4
	DESTROYED    int32 = 5
)

var PackSize int = 1400              // package size(default:MTU)
var ReadBufSize int = 64 * PackSize  // total read buf size
var WriteBufSize int = 64 * PackSize // totaol write buf size

var HeaderBytes int = 1024

var HandShakeTimeout time.Duration = time.Duration(3000) * time.Millisecond
var ReadTimeout time.Duration = time.Duration(3000) * time.Millisecond
var WriteTimeout time.Duration = time.Duration(3000) * time.Millisecond
