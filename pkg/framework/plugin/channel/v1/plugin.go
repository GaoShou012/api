package channel_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"framework/class/channel"
	"time"
)

func keyOfClientsOfChannel(topic string) string {
	return fmt.Sprintf("channel:client_list:%s", topic)
}
func keyOfTopicStream(topic string) string {
	return fmt.Sprintf("channel:stream:%s", topic)
}

func keyOfClientNotice(uuid string) string {
	return fmt.Sprintf("channel:client:notice:%s", uuid)
}

var _ channel.Channel = &plugin{}

type plugin struct {
	opts *Options
}

func (p *plugin) Clients(topic string) (channel.Clients, error) {
	res, err := p.opts.redisClient.HGetAll(context.TODO(), keyOfClientsOfChannel(topic)).Result()
	if err != nil {
		return nil, err
	}

	clients := make(map[string]time.Time)
	for key, val := range res {
		t, err := time.Parse(time.RFC3339, val)
		if err != nil {
			desc := fmt.Errorf("频道的客户端列表item解析失败,Topic=%s,val=%s,解析时间错误\n", topic, val)
			return nil, desc
		}
		clients[key] = t
	}

	return clients, nil
}

func (p *plugin) Publish(topic string, message []byte) (messageId string, err error) {
	var clients channel.Clients
	clients, err = p.Clients(topic)
	if err != nil {
		return
	}
	messageId, err = p.PublishWithClients(topic, message, clients)
	return
}

func (p *plugin) PublishWithClients(topic string, message []byte, clients channel.Clients) (messageId string, err error) {
	messageId, err = p.Publish(topic, message)
	if err != nil {
		return
	}

	var encode []byte
	e := &Event{
		T: channel.EventTypeNotice,
		D: map[string]string{"Topic": topic, "MsgId": messageId},
	}
	encode, err = json.Marshal(e)
	if err != nil {
		return
	}
	for uuid, _ := range clients {
		p.opts.Stream.Push(keyOfClientNotice(uuid), encode)
	}
	return
}

func (p *plugin) Subscribe(topic string, clientUUID string, timeout time.Duration) error {
	_, err := p.opts.redisClient.HSet(context.TODO(), keyOfClientsOfChannel(topic), clientUUID, time.Now().String()).Result()
	if err != nil {
		return err
	}
	return nil
}

func (p *plugin) UnSubscribe(topic string, clientUUID string) error {
	_, err := p.opts.redisClient.HDel(context.TODO(), keyOfClientsOfChannel(topic), clientUUID).Result()
	return err
}

func (p *plugin) Release(topic string) error {
	_, err := p.opts.redisClient.Del(context.TODO(), keyOfClientsOfChannel(topic)).Result()
	return err
}
