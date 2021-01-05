package cs

import "time"

type Session struct {
	Id        string
	State     uint64
	CreatedAt time.Time
}

func (s *Session) GetSessionId() string {
	return s.Id
}

// 是否已经关闭
func (s *Session) IsClose() bool {
	return false
}