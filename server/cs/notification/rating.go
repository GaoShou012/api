package notification

import "api/cs/message"

type SessionRating struct {
	SessionId string
	Rating    uint64
	Comment   string
}

func (n *SessionRating) GetMessageType() message.MsgType {
	return message.MsgTypeOperation
}
func (n *SessionRating) GetContentType() string {
	return "session.rating"
}
