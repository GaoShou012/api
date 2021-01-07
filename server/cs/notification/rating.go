package notification

import (
	"api/cs/event"
	"api/cs/message"
)

type SessionRating struct {
	SessionId string
	Rating    uint64
	Comment   string
}

func (n *SessionRating) GetMessageType() event.MsgType {
	return event.MsgTypeOperation
}
func (n *SessionRating) GetContentType() string {
	return "session.rating"
}
