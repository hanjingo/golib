package core

import (
	"fmt"
	"testing"
)

// go test -v dbuffer.go dbuffer_test.go -test.run TestRead
func TestRead(t *testing.T) {
	db := NewDBuffer()
	n, err := db.Read([]byte("hell work"))
	fmt.Printf("db.Read([]bytes("+"hello work"+"))= %d, %v", n, err)
}

// go test -v dbuffer.go dbuffer_test.go -test.run TestWrite
func TestWrite(t *testing.T) {
	db := NewDBuffer()
	n, err := db.Write([]byte("hell work"))
	fmt.Printf("db.Write([]bytes("+"hello work"+"))= %d, %v", n, err)
}
