package network

type SessionI interface {
	Run() error
	Read(arg []byte) (int, error)
	Write(args ...[]byte) (int, error)
	Close()
	Destroy()
}

type ServerI interface {
	Listen(addr string) error
	Close()
}

type ClientI interface {
	Dial(addr string) (SessionI, error)
}
