package meta

type SessionId string

type Session interface {
	GetId() SessionId
	GetSessionId() string
}

type SessionInfo interface {

}