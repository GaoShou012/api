package robot

type EventType int8

const (
	EventTypeNewSession = iota
	EventTypeVisitorMessage
)

type Event interface {
	GetType() EventType
	GetData() []byte
	GetSessionId() string
	GetMerchantCode() string
}

type event struct {
	T EventType
	D []byte
	S string
	M string
}

func (e *event) GetType() EventType {
	return e.T
}
func (e *event) GetData() []byte {
	return e.D
}
func (e *event) GetSessionId() string {
	return e.S
}
func (e *event) GetMerchantCode() string {
	return e.M
}
