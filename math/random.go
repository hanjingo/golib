package math

import (
	"math/rand"
	"time"
)

func Shuffle(in ...interface{}) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(in); i > 1; i-- {
		n := r.Intn(i) // 1~i
		index := len(in) - i
		// do swap
		temp := in[index]
		in[index] = in[n]
		in[n] = temp
	}
}
