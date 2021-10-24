package types

import (
	"fmt"
	"testing"
)

type SubClass struct {
	Id   int32
	Name string
}

type Class struct {
	Id   int32
	Idx  uint32
	Name string
	Sub  *SubClass
}

var src = &Class{
	Id:   1,
	Idx:  2,
	Name: "test",
	Sub: &SubClass{
		Id:   3,
		Name: "sub_test",
	},
}
var dst = &Class{}

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
