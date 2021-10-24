package core

import (
	"fmt"
	"testing"
)

// go test -v circle_queue.go circle_queue_test.go -test.run TestPush
func TestPush(*testing.T) {
	q := NewCircleQueue(5)
	fmt.Printf("q.Len()=%d\n", q.Len())

	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	q.Push(6)
	fmt.Printf("q.Len()=%d\n", q.Len())
}

// go test -v circle_queue.go circle_queue_test.go -test.run TestPop
func TestPop(*testing.T) {
	q := NewCircleQueue(5)
	fmt.Printf("q.Len()=%d\n", q.Len())

	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	q.Push(6)
	fmt.Printf("q.Len()=%d\n", q.Len())
	l := q.Len()
	for i := int64(0); i < l+1; i++ {
		fmt.Printf("i:%d, v:%v\n", i, q.Pop())
	}
}

// go test -v circle_queue.go circle_queue_test.go -test.run TestRange
func TestRange(*testing.T) {
	q := NewCircleQueue(5)
	fmt.Printf("q.Len()=%d\n", q.Len())

	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	q.Push(6)
	fmt.Printf("q.Len()=%d\n", q.Len())
	q.Range(func(i int64, v interface{}) {
		fmt.Printf("i=%d, v=%v\n", i, v)
	})
}
