package cs

import (
	"api/cs/event"
	"api/global"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Session struct {
	Id              string
	Enable          bool
	State           uint64
	CreatorId       uint64
	CreatorType     string
	CreatorUsername string
	CreatorNickname string
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

// 检查访客是否已经创建了会话
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

/*
	@params
	customerDevice 访客设备，PC，手机
	customerIp 访客IP
	customerLocation 访客地址
*/
func CustomerCreateSession(client Client, customerDevice uint64, customerIp string, customerLocation string) (*Session, error) {
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

	// 创建新会话
	session := &Session{
		Id:              uuid.NewV1().String(),
		Enable:          true,
		State:           0,
		CreatorId:       client.UserId,
		CreatorType:     client.UserType,
		CreatorUsername: client.Username,
		CreatorNickname: client.Nickname,
		CreatorIp:       customerIp,
		CreatorLocation: customerLocation,
		CreatedAt:       time.Now(),
	}
	if err := global.IM.CreateChannel(session); err != nil {
		return nil, err
	}

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
