package meta

type SessionState uint8

const (
	SessionStateR = iota
	SessionStateQueuing
	SessionStateServing
	SessionStateClosed
)

type Session interface {
	GetSessionId() string
}
