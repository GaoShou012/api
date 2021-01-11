package robot

type Event interface {
	Type() string
	Data() []byte
}
