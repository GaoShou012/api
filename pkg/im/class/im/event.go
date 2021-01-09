package im

type Event interface {
	Id() string
	Data() string
}
