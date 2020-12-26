package control

import (
	"cs/env"
	"cs/meta"
)

func CreateSession(session meta.Session, creator meta.Client) (uint64, error) {
	// 创建会话
	if err := env.Session.Create(session.GetSessionId(), session, creator); err != nil {
		return 0, err
	}

	// 客户关联会话
	if err := env.Client.SetSession(creator, session); err != nil {
		return 0, err
	}
	// 会话关联客户
	if err := env.Session.SetClient(session, creator); err != nil {
		return 0, err
	}


}

/*
	访客创建会话
	queueCode 队列编码，创建会话后，需要进行排队
	client 客户信息
	session 会话信息
*/
func CustomerCreateSession(queueCode string, client meta.Client, session meta.Session) (uint64, error) {
	// 会话

	// 保存会话信息
	if err := env.Session.SetInfo(session); err != nil {
		return 0, err
	}

	// 加入会话
	if err := JoinSession(session, client); err != nil {
		return 0, err
	}

	// 加入队列
	// 更新会话状态
	{
		num, err := JoinQueue(queueCode, session.GetSessionId())
		if err != nil {
			return 0, err
		}
		if err := env.Session.SetState(session.GetSessionId(), meta.SessionStateQueuing); err != nil {
			return 0, err
		}
		return num, nil
	}
}

/*
	退出排队
*/
func CustomerLeaveQueue(queueCode string, client meta.Client, session meta.Session) error {
	key := QueueName(queueCode)
	// 退出队列
	if err := env.Queue.Leave(key, session.GetSessionId()); err != nil {
		return err
	}

	// 退出会话
	if err := LeaveSession(session, client); err != nil {
		return err
	}

	// 更新会话状态
	if err := env.Session.SetState(session.GetSessionId(), meta.SessionStateClosed); err != nil {
		return err
	}

	return nil
}
