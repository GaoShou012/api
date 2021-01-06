package im_v1

type syncRecordEventType int

const (
	syncRecordEventTypeClientAttach syncRecordEventType = iota
	syncRecordEventTypeClientDetach
)

type syncRecordEventClientAttach struct {
	UUID          string
	LastMessageId string
}
type syncRecordEventClientDetach struct {
	UUID string
}
type syncRecordEvent struct {
	Type syncRecordEventType
	Data interface{}
}
