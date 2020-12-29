package control

import (
	"cs/env"
	"cs/meta"
)

func CreateSession(creator meta.Client, session meta.Session, ) error {
	// 创建会话
	if err := env.Session.Create(session, creator); err != nil {
		return err
	}
	return nil
}

func JoinSession(client meta.Client, session meta.Session) error {
	// 客户关联会话
	if err := env.Client.SetSession(client, session); err != nil {
		return err
	}
	// 会话关联客户
	if err := env.Session.SetClient(session, client); err != nil {
		return err
	}
	return nil
}

func LeaveSession(session meta.Session, client meta.Client) error {
	// 客户取消关联会话
	if err := env.Client.DelSession(client, session); err != nil {
		return err
	}
	// 会话取消关联客户
	if err := env.Session.DelClient(session, client); err != nil {
		return err
	}
	return nil
}

// 是否有加入会话
func ExistsSession(client meta.Client, sessionId string) (bool, error) {
	return env.Client.ExistsSession(client, sessionId)
}

func Broadcast(session meta.Session, data []byte) error {
	clients, err := env.Session.GetAllClients(session)
	if err != nil {
		return err
	}
	for _, client := range clients {
		if err := env.Gateway.Publish(client.GetUUID(), data); err != nil {
			return err
		}
	}
	return nil
}
