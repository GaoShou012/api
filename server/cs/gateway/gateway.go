package gateway

import (
	"api/cs"
	"api/cs/env"
	"api/global"
	"api/libs/watcher"
	"context"
	"fmt"
)

type Gateway struct {
}

func (g *Gateway) Send() {

}

type ClientConn interface {
	Publish(message []byte) error
	Close() error
}

type MessageCache struct {
	Cache chan []byte
	context.CancelFunc
}

func (i *MessageCache) IsUpLimit() bool {
	if len(i.Cache) >= 200 {
		return true
	}
	return false
}
func (i *MessageCache) UpLimitCallback() {
	i.CancelFunc()
}

type Client struct {
	UUID          string
	LastMessageId string

	Events                chan *cs.Event
	Publish               chan []byte
	SendCount             uint64
	Conn                  ClientConn
	Exit                  context.CancelFunc
	MessageCacheOnSending chan []byte
}

func (c *Client) IsUpLimit() bool {
	if len(c.MessageCacheOnSending) >= 64 {
		return true
	}
	return false
}
func (c *Client) UpLimitCallback() {
	c.Exit()
}

func (c *Client) Init() {
	c.MessageCacheOnSending = make(chan []byte, 256)

	overload, cancel := context.WithCancel(context.Background())
	c.Exit = cancel

	watcher.Add(c)

	go func() {
		for {
			select {
			case <-overload.Done():
				global.Logger.Info(fmt.Sprintf("UUID:%s关闭通知channel", c.UUID))
				return
			case event := <-c.Events:
				// 收到通知，进行拉取会话消息
				// 消息投递到发送器
				switch event.Type {
				case cs.EventTypeNotification:
					message, err := env.Session.GetMessageById(event.Topic, event.MsgId)
					if err != nil {
						global.Logger.Error(err)
						continue
					}
					c.MessageCacheOnSending <- message
					break
				}
			}
		}
	}()
	go func() {
		for {
			select {
			case <-overload.Done():
				global.Logger.Info(fmt.Sprintf("UUID:%s,关闭推送程序", c.UUID))
				c.Conn.Close()
				return
			case message := <-c.Publish:
				c.Conn.Publish(message)
			}
		}
	}()
}
