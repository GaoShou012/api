package channel

type Event interface {
	Id() string
	Data() []byte
}
