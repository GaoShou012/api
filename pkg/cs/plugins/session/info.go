package session

import (
	"cs/class/session"
)

var _ session.Info = &Info{}

type Info struct {
	Enable bool
	Info   []byte
}

func (i Info) GetId() string {
	panic("implement me")
}

func (i Info) GetKey() string {
	panic("implement me")
}
