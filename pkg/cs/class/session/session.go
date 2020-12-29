package session

import "cs/meta"

type Session interface {
	Init() error

	// 创建会话
	// client是创建者
	Create(session meta.Session, creator meta.Client) error

	Delete(session meta.Session) error

	// 创建者信息
	GetCreator(sessionId string) (meta.Client, error)

	// 设置会话信息
	SetInfo(sessionId string, session meta.Session) error
	// 获取会话信息
	GetInfo(sessionId string) (meta.Session, error)

	// 查找会话信息
	Find(sessionId string,session meta.Session) (bool,error)

	// 设置标记
	SetFlag(sessionId string, key string, val string) error
	SetFlagNX(sessionId string, key string, val string) (bool, error)
	GetFlag(sessionId string, key string) (string, error)
	DelFlag(sessionId string, key string) error
	ExistsFlag(sessionId string, key string) (bool, error)

	SetEnable(sessionId string, enable bool) (bool, error)
	// 会话是否启用中
	GetEnable(sessionId string) (bool, error)

	SetState(sessionId string, state meta.SessionState) error
	GetState(sessionId string) (meta.SessionState, error)

	// 保存会话信息
	//SetInfo(session meta.Session) error
	// 读取会话信息
	//GetInfo(sessionId string) (meta.Session, error)
	// 会话信息是否存在
	ExistsInfo(sessionId string) (bool, error)
	// 移除会话信息
	DelInfo(session meta.Session) error

	// 序列化会话信息
	// 可在应用层，扩展真实数据
	MarshalSessionInfo(session meta.Session) error
	// 反序列会话信息
	// 可在应用层，扩展真实数据
	UnmarshalSessionInfo(data []byte) (meta.Session, error)

	// 增加客户
	SetClient(session meta.Session, client meta.Client) error
	// 移除客户
	DelClient(session meta.Session, client meta.Client) error
	// 客户是否存在
	ExistsClient(session meta.Session, client meta.Client) (bool, error)
	// 客户列表
	GetAllClients(session meta.Session) ([]meta.Client, error)

	// 移除整个客户端列表
	DelAllClients(session meta.Session) error

	// 压入消息
	PushMessage(session meta.Session, message []byte) (string, error)
	// 拉取消息
	PullMessage(session meta.Session, lastMessageId string, count uint64) ([][]byte, error)
}
