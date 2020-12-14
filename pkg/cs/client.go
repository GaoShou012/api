package cs

import (
	"cs/global"
	"cs/meta"
	"errors"
)

type ClientInfo struct {
}

type Session struct {
}

func CreateSession(session meta.Session) error {
	if err := Adapter.SaveSession(session); err != nil {
		return err
	}
	return nil
}

/*
	加入会话
*/
func JoinSession(client meta.Client, session meta.Session, ) error {
	// 会话增加客户ID
	{
		ok, err := global.Session.AddClient(client)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("会话增加客户失败")
		}
	}

	// 客户增加会话ID
	{
		ok, err := global.Client.AddSession(session.GetId(), session)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("客户增加会话失败")
		}
	}

	return nil
}

/*
	离开会话
*/
func LeaveSession(client meta.Client, sessionId meta.SessionId) error {
	// 会话是否存在此客户
	{
		ok, err := global.Session.ExistsClient(client)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("客户不存在此会话中")
		}
	}

	// 会话移除客户ID
	{
		ok, err := global.Session.DelClient(client)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("会话移除客户失败")
		}
	}

	// 客户移除会话ID
	{
		ok, err := global.Client.DelSession(sessionId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("客户移除会话失败")
		}
	}

	return nil
}
