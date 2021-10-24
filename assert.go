package golib

func Assert(condition bool, desc string) {
	if !condition {
		panic(desc)
	}
}
