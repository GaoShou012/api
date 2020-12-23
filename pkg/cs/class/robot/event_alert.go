package robot

type EventType uint8

const (
	// 建立会话
	EventTypeEstablishSession EventType = iota
	// 客户说话
	EventTypeCustomerTalk
	// 客服说话
	EventTypeCsTalk
)

func (e EventType) String() string {
	switch e {
	case EventTypeEstablishSession:
		return "建立会话"
	case EventTypeCustomerTalk:
		return "访客说话"
	case EventTypeCsTalk:
		return "客服说话"
	default:
		return "未知"
	}
}

/*
	提醒事件
*/
type AlertEvent interface {
	GetSessionId() string
	GetEventType() EventType
}
