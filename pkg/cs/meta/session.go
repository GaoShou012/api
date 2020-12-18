package meta

type SessionState uint8

const (
	SessionStateR = iota
	SessionStateQueuing
	SessionStateServing
	SessionStateClosed
)

type Session interface {
	NewInfo(client Client)
	GetSessionId() string
}