package cs

import (
	"api/cs/env"
	"api/cs/event"
	"api/cs/meta"
	"api/global"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Session struct {
	Id              string
	Enable          bool
	State           uint64
	Creator         *meta.Client
	CreatorIp       string
	CreatorLocation string
	CreatedAt       time.Time
}

func (s *Session) GetTopic() string {
	return s.Id
}
func (s *Session) GetEnable() bool {
	return s.Enable
}
func (s *Session) SetEnable(enable bool) {
	s.Enable = enable
}

func topicOfSessionChannel(sessionId string) string {
	return fmt.Sprintf("session:channel:%s", sessionId)
}

// 检查访客是否已经创建了会话
// 如果已经创建，返回会话信息
func GetCustomerSession(clientUUID string) (*Session, error) {
	channels, err := global.IM.Client().Channels(clientUUID)
	if err != nil {
		return nil, err
	}

	if len(channels) == 0 {
		return nil, nil
	}

	var topic string
	for key, _ := range channels {
		topic = key
	}
	session := &Session{}
	// 读取会话信息
	if err := global.IM.Channel().GetInfo(topic, session); err != nil {
		return nil, err
	}
	return session, nil
}

// 保存会话信息
func GetSessionInfo(sessionId string) (*Session, error) {
	topic := topicOfSessionChannel(sessionId)
	session := &Session{}
	if err := global.IM.Channel().GetInfo(topic, session); err != nil {
		return nil, err
	}
	return session, nil
}

// 访客、客服，通过此方法，发送消息
func SessionMessage(sessionId string, sender Sender, content string, contentType string) (string, error) {
	topic := topicOfSessionChannel(sessionId)

	data, err := event.Encode(&event.ClientMessage{
		Sender:      GetSenderInfoFrom(sender),
		Content:     content,
		ContentType: contentType,
		Time:        time.Now(),
	})
	if err != nil {
		return "", err
	}
	messageId, err := global.IM.Channel().Push(topic, data)
	if err != nil {
		return "", err
	}

	return messageId, nil
}

/*
	@params
	customerDevice 访客设备，PC，手机
	customerIp 访客IP
	customerLocation 访客地址
*/
func CustomerCreateSession(client *meta.Client, customerDevice uint64, customerIp string, customerLocation string) (*Session, error) {
	merchantCode := client.MerchantCode
	clientUUID := client.GetUUID()

	{
		// TODO 访客分布式操作锁
	}

	if session, err := GetCustomerSession(clientUUID); err != nil {
		return nil, err
	} else {
		if session != nil {
			return session, nil
		}
	}

	// 会话ID
	sessionId := uuid.NewV1().String()
	timestamp := time.Now().Unix()

	// 先把会话ID，存储到会话集合中，用于异步对会话进行回收
	if err := env.SessionsSet.SetItem(TopicOfSessionsSet, sessionId, float64(timestamp)); err != nil {
		return nil, err
	}

	// 先把会话ID，存储到商户的会话集合中，用于统计会话信息
	if err := env.SessionsSet.SetItem(TopicOfMerchantSessionSet(merchantCode), sessionId, float64(timestamp)); err != nil {
		return nil, err
	}

	// 创建新会话
	session := &Session{
		Id:              sessionId,
		Enable:          true,
		State:           0,
		Creator:         client,
		CreatorIp:       customerIp,
		CreatorLocation: customerLocation,
		CreatedAt:       time.Now(),
	}
	if err := global.IM.CreateChannel(session); err != nil {
		return nil, err
	}

	// 访客订阅此会话频道
	topic := session.GetTopic()
	if err := global.IM.ClientSubscribeChannel(clientUUID, topic); err != nil {
		return nil, err
	}

	// TODO 加入队列，分配会话给客服

	go func() {
		// 统计在线会话数量，广播给所有客服
		{
			uuids := make([]string, 0)
			num, err := statsSessionOpening()
			if err != nil {
				global.Logger.Error(err)
				return
			}
			data, err := event.Encode(&event.MonitorTotalSessions{Data: num})
			if err != nil {
				global.Logger.Error(err)
				return
			}
			for _, id := range uuids {
				if err := global.IM.PushMessageToClient(id, data); err != nil {
					global.Logger.Error(err)
					continue
				}
			}
		}

		// 会话信息入库
		{
			// TODO model
		}

		// 访客信息入库
		{
			// TODO model
		}
	}()

	return session, nil
}
