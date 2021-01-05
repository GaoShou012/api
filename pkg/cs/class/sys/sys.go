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

	// 广播消息
	// 消息存储到stream
	// 然后广播通知到各个私人stream
	// 并且异步消息入库
	Broadcast(client meta.Client, sessionId string, data interface{}) (string, error)
}
