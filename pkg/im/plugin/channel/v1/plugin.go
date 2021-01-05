package v1

import (
	"encoding/json"
	"fmt"
	"im/class/channel"
	"im/env"
	"time"
)

// 频道信息
func keyOfChannelInfo(topic string) string {
	return fmt.Sprintf("channel:info:%s", topic)
}

// 频道消息流
func keyOfChannelMessage(topic string) string {
	return fmt.Sprintf("channel:message:%s", topic)
}

// 频道客户端列表
func keyOfClientListOfChannel(topic string) string {
	return fmt.Sprintf("channel:client_list:%s", topic)
}

var _ channel.Channel = &plugin{}

type plugin struct {
	opts *Options
}

// 如果不存在返回频道信息不存在
func (p *plugin) isInfoExists(topic string) error {
	num, err := p.opts.redisClient.Exists(keyOfChannelInfo(topic)).Result()
	if err != nil {
		return err
	}
	if num == 0 {
		return fmt.Errorf("频道信息不存在")
	}
	if num == 1 {
		return nil
	} else {
		return fmt.Errorf("频道信息key(%s),不是唯一性", topic)
	}
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) Create(info channel.Info) error {
	return p.SetInfo(info.GetTopic(), info)
}

func (p *plugin) Delete(topic string) error {
	// 删除频道信息
	{
		_, err := p.opts.redisClient.Del(keyOfChannelInfo(topic)).Result()
		if err != nil {
			return err
		}
	}
	// 删除消息流
	{
		_, err := p.opts.redisClient.Del(keyOfChannelMessage(topic)).Result()
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *plugin) SetEnable(topic string, enable bool) error {
	if err := p.isInfoExists(topic); err != nil {
		return err
	}

	_, err := p.opts.redisClient.HSet(keyOfChannelInfo(topic), "Enable", enable).Result()
	if err != nil {
		return err
	}
	return nil
}

func (p *plugin) GetEnable(topic string) bool {
	res, err := p.opts.redisClient.HGet(keyOfChannelInfo(topic), "Enable").Result()
	if err != nil {
		env.Logger.Error(err)
		return false
	}
	if res == "1" {
		return true
	} else {
		return false
	}
}

func (p *plugin) SetInfo(topic string, info channel.Info) error {
	if err := p.isInfoExists(topic); err != nil {
		return err
	}

	j, err := json.Marshal(info)
	if err != nil {
		return err
	}
	m := make(map[string]interface{})
	m["Enable"] = info.GetEnable()
	m["Info"] = j

	_, err = p.opts.redisClient.HMSet(keyOfChannelInfo(topic), m).Result()
	if err != nil {
		return err
	}
	return nil
}

func (p *plugin) GetInfo(topic string, info channel.Info) error {
	if err := p.isInfoExists(topic); err != nil {
		return err
	}
	res, err := p.opts.redisClient.HGetAll(keyOfChannelInfo(topic)).Result()
	if err != nil {
		return err
	}
	data := []byte(res["Info"])
	if err := json.Unmarshal(data, info); err != nil {
		return err
	}
	if res["Enable"] == "1" {
		info.SetEnable(true)
	} else {
		info.SetEnable(false)
	}
	return nil
}

func (p *plugin) Clients(topic string) (channel.Clients, error) {
	res, err := p.opts.redisClient.HGetAll(keyOfClientListOfChannel(topic)).Result()
	if err != nil {
		return nil, err
	}

	clients := make(channel.Clients)
	for key, val := range res {
		t, err := time.Parse(time.RFC3339, val)
		if err != nil {
			return nil, err
		}
		clients[key] = t
	}

	return clients, nil
}

func (p *plugin) Publish(topic string, message []byte) (messageId string, err error) {
	// 检查频道是否开启
	if p.GetEnable(topic) == false {
		err = fmt.Errorf("频道未启用，不能推送消息")
		return
	}

	// 消息推入频道
	messageId, err = p.opts.Stream.Push(topic, message)
	if err != nil {
		return
	}

	return
}

func (p *plugin) Subscribe(topic string, clientUUID string) error {
	// 检查频道是否开启
	if p.GetEnable(topic) == false {
		return fmt.Errorf("频道未启用，不能操作")
	}

	_, err := p.opts.redisClient.HSet(keyOfClientListOfChannel(topic), clientUUID, time.Now().String()).Result()
	return err
}

func (p *plugin) UnSubscribe(topic string, clientUUID string) error {
	// 检查频道是否开启
	if p.GetEnable(topic) == false {
		return fmt.Errorf("频道未启用，不能操作")
	}

	_, err := p.opts.redisClient.HDel(keyOfClientListOfChannel(topic), clientUUID).Result()
	return err
}

func (p *plugin) Release(topic string) error {
	_, err := p.opts.redisClient.Del(keyOfClientListOfChannel(topic)).Result()
	return err
}
