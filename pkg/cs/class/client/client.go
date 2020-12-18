package client

import "cs/meta"

type Client interface {
	Init() error

	// 保存客户信息
	SetInfo(uuid string, client meta.Client) error
	// 获取客户信息
	GetInfo(uuid string, client meta.Client) (bool, error)
	// 客户信息是否存在
	ExistsInfo(uuid string) (bool, error)

	// 客户会话列表，加入会话
	SetSession(client meta.Client, sessionId string, session meta.Session) error
	// 客户会话列表，移除会话
	DelSession(client meta.Client, sessionId string) error
	// 客户会话列表，会话是否存在
	ExistsSession(client meta.Client, sessionId string) (bool, error)
	// 获取客户会话列表
	GetAllSessions(client meta.Client) ([]string, error)
	// 加入的会话数量
	GetNumSessions(client meta.Client) (uint64, error)
}
