package im_v1

type clientEvent struct {
	id   string
	data []byte
}

func (e *clientEvent) Id() string {
	return e.id
}
func (e *clientEvent) Data() []byte {
	return e.data
}

const (
	eventTypeMessage = iota
	eventTypeChannelNotice
)

type clientEventOfChannelNotice struct {
	Topic     string
	MessageId string
}
