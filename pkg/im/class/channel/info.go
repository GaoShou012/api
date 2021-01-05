package channel

type Info interface {
	GetTopic() string
	SetEnable(enable bool)
	GetEnable() bool
}
