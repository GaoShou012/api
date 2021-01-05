package channel_v1

import (
	"context"
	"cs/class/channel"
	"cs/class/client"
	"cs/env"
	"encoding/json"
	"fmt"
	"time"
)

func keyOfClientsOfChannel(topic string) string {
	return fmt.Sprintf("channel:%s:clients", topic)
}

var _ channel.Channel = &plugin{}

type plugin struct {
	opts *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) Publish(topic string, message []byte) (messageId string, err error) {
	msgId, err := p.opts.Stream.Push(topic, message)

	// 读取所有订阅者
	res, err := p.opts.redisClient.HGetAll(context.TODO(), keyOfClientsOfChannel(topic)).Result()
	if err != nil {
		desc := fmt.Sprintf("获取所有订阅的客户失败,topic=%s,err=%v", topic, err)
		env.Logger.Error(desc)
		return msgId, nil
	}

	// 遍历订阅的客户端列表
	// 推送通知
	for k, v := range res {
		cli := &channel.Client{}
		if err := json.Unmarshal([]byte(v), cli); err != nil {
			desc := fmt.Sprintf("解析client item失败,topic=%s,k=%s,v=%s,err=%v", topic, k, v, err)
			env.Logger.Error(desc)
			continue
		}
		event := &client.Event{
			Type: client.EventTypeNotice,
			EventNotice: &client.EventNotice{
				Topic: topic,
				MsgId: msgId,
			},
			EventMessage: nil,
		}
		if err := env.Client.PushEvent(cli.UUID, event); err != nil {
			desc := fmt.Sprintf("推送通知失败,uuid=%s,err=%v", cli.UUID, err)
			env.Logger.Error(desc)
			continue
		}
	}

	return msgId, nil
}

func (p *plugin) Subscribe(topic string, uuid string, timeout time.Duration) error {
	item := &channel.Client{
		UUID:      uuid,
		Timeout:   timeout,
		CreatedAt: time.Now(),
	}
	j, err := json.Marshal(item)
	if err != nil {
		return err
	}
	_, err = p.opts.redisClient.HSet(context.TODO(), keyOfClientsOfChannel(topic), j).Result()
	return err
}

func (p *plugin) UnSubscribe(topic string, uuid string) error {
	_, err := p.opts.redisClient.HDel(context.TODO(), keyOfClientsOfChannel(topic), uuid).Result()
	return err
}

func (p *plugin) Release(topic string) error {
	_, err := p.opts.redisClient.Del(context.TODO(), keyOfClientsOfChannel(topic)).Result()
	return err
}
