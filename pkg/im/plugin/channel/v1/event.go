package channel_v1

type event struct {
	msgId   string
	msgData []byte
}

func (e *event) Id() string {
	return e.msgId
}
func (e *event) Data() []byte {
	return e.msgData
}
