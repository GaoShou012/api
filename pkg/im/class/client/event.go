package client

type Event interface {
	Id() string
	Data() []byte
}
