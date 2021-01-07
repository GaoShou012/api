package notification

type SessionRating struct {
	SessionId string
	Rating    uint64
	Comment   string
}

func (n *SessionRating) GetType() string {
	return "session.rating"
}
