package core

import (
	"bytes"
	"sync"
)

type DBuffer struct {
	mu   *sync.RWMutex
	rbuf *bytes.Buffer
	wbuf *bytes.Buffer
}

func NewDBuffer() *DBuffer {
	return &DBuffer{
		mu:   new(sync.RWMutex),
		rbuf: bytes.NewBuffer([]byte{}),
		wbuf: bytes.NewBuffer([]byte{}),
	}
}

func (d *DBuffer) Read(data []byte) (int, error) {
	if d.rbuf.Len() == 0 {
		d.Switch()
	}
	return d.rbuf.Read(data)
}

func (d *DBuffer) Write(data []byte) (int, error) {
	return d.wbuf.Write(data)
}

func (d *DBuffer) Switch() {
	d.mu.Lock()
	defer d.mu.Unlock()

	tmp := d.wbuf
	d.wbuf = d.rbuf
	d.rbuf = tmp
}
