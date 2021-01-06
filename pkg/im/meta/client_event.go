package meta

type ClientEventType int

const (
	ClientEventTypeChannel ClientEventType = iota
	ClientEventTypeMessage
)

type ClientEventChannel struct {
	// 目标频道名称
	Topic string
	// 目标频道消息ID
	MsgId string
}
type ClientEventMessage struct {
	Payload []byte
}

type ClientEvent struct {
	Type ClientEventType
	*ClientEventChannel
	*ClientEventMessage
}
