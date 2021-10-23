package golib

import (
	"fmt"
	"testing"
)

type sub_class struct {
	id   int32
	name string
}

type class struct {
	id   int32
	idx  uint32
	name string
	sub  *sub_class
}

var src = &class{
	id:   1,
	idx:  2,
	name: "test",
	sub: &sub_class{
		id:   3,
		name: "sub_test",
	},
}
var dst = &class{}

// go test -v clone.go clone_test.go -test.run TestDeepCopy
func TestDeepCopy(t *testing.T) {
	DeepCopy(dst, src)
	fmt.Printf("DeepCopy(dst, src), src = %v, dts = %v", src, dst)
}

// go test -v clone.go clone_test.go -test.run TestDeepClone
func TestDeepClone(t *testing.T) {
	dst1 := DeepClone(src)
	fmt.Printf("DeepCopy(src), src = %v, dts1 = %v", src, dst1)
}