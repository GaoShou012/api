package robot

import "encoding/json"

type EventType uint16

const (
	EventTypeNewSession EventType = iota
	EventTypeVisitorMessage
)

type Event interface {
	MerchantCode() string
	SessionId() string
	Type() EventType
	Data() []byte
}

type event struct {
	t EventType
	d []byte
}

func (e *event) Type() EventType {
	return e.t
}
func (e *event) Data() []byte {
	return e.d
}

func encodeEvent(event Event) ([]byte, error) {
	j, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	data := make([]byte, len(j)+2)

	t := event.Type()
	tH := uint8(t >> 8 & 0xff)
	tL := uint8(t & 0xff)
	data[0] = tH
	data[1] = tL
	for i := 0; i < len(j); i++ {
		data[i+2] = j[i]
	}
	return data, nil
}
func decodeEvent(data []byte) (Event, error) {
	tH := uint16(data[0]) << 8 & 0xff00
	tL := uint16(data[1]) & 0x00ff
	evt := &event{
		t: EventType(tH & tL),
		d: data[2:],
	}
	return evt, nil
}

// 新会话事件
type EventOfNewSession struct {
	MerchantCode string
	SessionId    string
	ClientUUID   string
}

func (e *EventOfNewSession) Type() EventType {
	return EventTypeNewSession
}

// 访客消息
type EventOfCustomerMessage struct {
	MerchantCode string
	SessionId    string
	ClientUUID   string
	Content      string
	ContentType  string
}
