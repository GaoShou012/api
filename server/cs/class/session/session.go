package session

type Session interface {
	GetMessageById(sessionId string, messageId string) ([]byte, error)
	Broadcast(sessionId string, message []byte) error
}
