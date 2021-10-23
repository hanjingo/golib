package golib

type SIZE int64

const (
	BYTE SIZE = 1
	KB   SIZE = BYTE << 10
	MB   SIZE = KB << 10
	GB   SIZE = MB << 10
	TB   SIZE = GB << 10
)

func (s SIZE) Add(arg SIZE) SIZE {
	return s + arg
}

func (s SIZE) Del(arg SIZE) SIZE {
	return s - arg
}

func (s SIZE) TB() SIZE {
	return s / TB
}

func (s SIZE) GB() SIZE {
	return s / GB
}

func (s SIZE) MB() SIZE {
	return s / MB
}

func (s SIZE) KB() SIZE {
	return s / KB
}
