package session

import "cs/meta"

type Session interface {
	Init() error
	SetEnable(sessionId string, enable bool) (bool, error)
	// 会话是否启用中
	GetEnable(sessionId string) (bool, error)

	// 保存会话信息
	SetInfo(session meta.Session) error
	// 读取会话信息
	GetInfo(sessionId string, session meta.Session) (bool, error)
	// 会话信息是否存在
	ExistsInfo(sessionId string) (bool, error)
	// 移除会话信息
	DelInfo(session meta.Session) error

	// 增加客户
	SetClient(session meta.Session, client meta.Client) error
	// 移除客户
	DelClient(session meta.Session, client meta.Client) error
	// 客户是否存在
	ExistsClient(session meta.Session, client meta.Client) (bool, error)
	// 客户列表
	GetAllClients(session meta.Session, clients interface{}) error
	// 移除整个客户端列表
	DelAllClients(session meta.Session) error

	// 压入消息
	PushMessage(session meta.Session, message []byte) (string, error)
	// 拉取消息
	PullMessage(session meta.Session, lastMessageId string, count uint64) ([][]byte, error)
}
