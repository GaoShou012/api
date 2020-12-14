package class

import "cs/meta"

type SessionId string

type Client interface {

	// 缓存客户信息
	//SaveInfo(client ClientInfo) error
	// 读取缓存信息
	//ReadInfo(uuid string, client ClientInfo) error

	// 加入会话
	AddSession(id meta.SessionId, sessionInfo meta.Session) (bool, error)

	// 离开会话
	DelSession(id meta.SessionId) (bool, error)

	// 是否已经加入会话
	ExistsSession(id meta.SessionId) (bool, error)
}

type ClientInfo interface {
	GetTenantId() uint64
	GetTenantCode() string
	GetUserType() string
	GetUserId() string
}

type ClientSession interface {
	// 加入会话
	Add(id SessionId) (bool, error)
	// 离开会话
	Del(id SessionId) (bool, error)
	// 是否存在会话
	Exists(id SessionId) (bool, error)
	// 会话数量
	Len() (uint64, error)
}
