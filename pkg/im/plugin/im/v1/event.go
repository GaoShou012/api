package im_v1

type clientEventType int

const (
	clientEventTypeNotice clientEventType = iota
	clientEventTypeMessage
)

type clientEvent struct {
	Type  clientEventType
	Data  []byte
	MsgId string
	Topic string
}
