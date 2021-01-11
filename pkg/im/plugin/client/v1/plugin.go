package client_v1

import (
	"fmt"
	"im/class/client"
	"im/env"
	"time"
)

func keyOfClientMessageStream(uuid string) string {
	return fmt.Sprintf("client:message:%s", uuid)
}

func keyOfClientChannels(uuid string) string {
	return fmt.Sprintf("client:channels:%s", uuid)
}

var _ client.Client = &plugin{}

type plugin struct {
	opts *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) Push(uuid string, message []byte) (messageId string, err error) {
	return p.opts.Stream.Push(keyOfClientMessageStream(uuid), message)
}

func (p *plugin) PushClients(uuid []string, message []byte) error {
	for _, id := range uuid {
		if _, err := p.opts.Stream.Push(keyOfClientMessageStream(id), message); err != nil {
			env.Logger.Error(err)
			continue
		}
	}
	return nil
}

func (p *plugin) Pull(uuid string, lastMessageId string, count uint64) ([]client.Event, error) {
	res, err := p.opts.Stream.Pull(keyOfClientMessageStream(uuid), lastMessageId, count)
	if err != nil {
		return nil, err
	}

	i := 0
	events := make([]client.Event, len(res))

	for _, val := range res {
		evt := &event{
			msgId:   val.Id(),
			msgData: val.Message(),
		}
		events[i] = evt
		i++
	}

	return events, nil
}

func (p *plugin) PullById(uuid string, messageId string) (client.Event, error) {
	res, err := p.opts.Stream.PullById(keyOfClientMessageStream(uuid), messageId)
	if err != nil {
		return nil, err
	}
	evt := &event{
		msgId:   res.Id(),
		msgData: res.Message(),
	}
	return evt, nil
}

func (p *plugin) Delete(uuid string) error {
	return p.opts.Stream.Delete(keyOfClientMessageStream(uuid))
}

func (p *plugin) Subscribe(uuid string, topic string) error {
	_, err := p.opts.redisClient.HSet(keyOfClientChannels(uuid), topic, time.Now().String()).Result()
	return err
}

func (p *plugin) UnSubscribe(uuid string, topic string) error {
	_, err := p.opts.redisClient.HDel(keyOfClientChannels(uuid), topic).Result()
	return err
}

func (p *plugin) Channels(uuid string) (client.Channels, error) {
	res, err := p.opts.redisClient.HGetAll(keyOfClientChannels(uuid)).Result()
	if err != nil {
		return nil, err
	}
	channels := make(client.Channels)
	for key, val := range res {
		t, err := time.Parse(time.RFC3339, val)
		if err != nil {
			return nil, err
		}
		channels[key] = t
	}
	return channels, nil
}
