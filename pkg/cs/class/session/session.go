package session

import "cs/meta"

type Session interface {
	SetEnable(sessionId string, enable bool) (bool, error)
	// 会话是否启用中
	GetEnable(sessionId string) (bool, error)

	// 保存会话信息
	SaveInfo(sessionId string, info Info) error
	// 读取会话信息
	ReadInfo(sessionId string, info Info) (bool, error)
	// 会话信息是否存在
	ExistsInfo(sessionId string) (bool, error)

	// 增加客户
	AddClient(sessionId string, client meta.Client) (bool, error)
	// 移除客户
	DelClient(sessionId string, client meta.Client) (bool, error)
	// 客户是否存在
	ExistsClient(sessionId string, client meta.Client) (bool, error)
}
