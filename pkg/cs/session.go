package cs

import (
	"cs/env"
	"cs/meta"
)

/*
	加入会话
*/
func JoinSession(session meta.Session, client meta.Client) error {
	if err := env.Session.SetClient(session, client); err != nil {
		return err
	}

	if err := env.Client.SetSession(client, session.GetSessionId(), session); err != nil {
		return err
	}

	return nil
}

/*
	离开会话
*/
func LeaveSession(session meta.Session, client meta.Client) error {
	if err := env.Session.DelClient(session, client); err != nil {
		return err
	}

	if err := env.Client.DelSession(client, session.GetSessionId()); err != nil {
		return err
	}

	return nil
}

/*
	关闭会话
	@params
	session 会话
	clients 会话-所有成员
*/
func CloseSession(session meta.Session, clients []meta.Client) error {
	// 移除会话信息
	if err := env.Session.DelInfo(session); err != nil {
		return err
	}

	// 移除会话成员列表
	if err := env.Session.DelAllClients(session); err != nil {
		return err
	}

	// 从各个成员的会话列表中移除
	for _, client := range clients {
		if err := env.Client.DelSession(client, session.GetSessionId()); err != nil {
			return err
		}
	}

	return nil
}

/*
	会话广播消息
*/
func Broadcast(session meta.Session, clients []meta.Client, message []byte) error {
	//for _, client := range clients {
	//
	//}
	return nil
}
