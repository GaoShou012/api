package session

type Info interface {
	GetId() string
	GetKey() string
}
