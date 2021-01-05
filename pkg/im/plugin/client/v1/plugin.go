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

func (p *plugin) Push(uuid string, eventData client.EventData) error {
	evt, err := EncodeEvent(eventData)
	if err != nil {
		return err
	}
	_, err = p.opts.Stream.Push(keyOfClientMessageStream(uuid), evt)
	return err
}

func (p *plugin) PushClients(uuid []string, eventData client.EventData) error {
	evt, err := EncodeEvent(eventData)
	if err != nil {
		return err
	}
	for _, id := range uuid {
		if _, err := p.opts.Stream.Push(keyOfClientMessageStream(id), evt); err != nil {
			env.Logger.Error(err)
			continue
		}
	}
	return nil
}

func (p *plugin) Pull(uuid string, count int) ([]client.Event, error) {
	panic("implement me")
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
