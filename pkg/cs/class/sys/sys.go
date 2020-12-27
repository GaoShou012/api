package sys

import "cs/meta"

type Sys interface {
	// 创建会话
	CreateSession(client meta.Client, session meta.Session) error
	// 加入会话
	JoinSession(client meta.Client, session meta.Session) error
	// 离开会话
	LeaveSession(client meta.Client, session meta.Session) error
	// 客户是否加入了会话
	IsClientInSession(client meta.Client, session meta.Session) (bool, error)
}
