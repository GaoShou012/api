package control

import (
	"cs/env"
	"cs/meta"
	"fmt"
)

/*
	客服接待会话
	queueCode 队列编码
	client 需要分配的客服
*/
func CsReceivingSession(queueCode string, client meta.Client) (bool, error) {
	key := QueueName(queueCode)

	// 判断队列是否存在会话
	{
		num, err := env.Queue.Len(key)
		if err != nil {
			return false, err
		}
		if num <= 0 {
			return false, nil
		}
	}

	// 获取队列最前的一个会话
	item, err := env.Queue.GetTheFirstSession(key)
	if err != nil {
		return false, err
	}

	// 会话ID
	sessionId := item.SessionId

	// 读取会话信息
	session, err := env.Session.GetInfo(sessionId)
	if err != nil {
		return false, err
	}
	if session == nil {
		return false, fmt.Errorf("会话信息不存在,会话ID:%s", sessionId)
	}

	// 加入会话
	// 更新状态
	{
		if err := JoinSession(session, client); err != nil {
			return false, err
		}
		if err := env.Session.SetState(sessionId, meta.SessionStateServing); err != nil {
			return false, nil
		}
	}

	return true, nil
}

