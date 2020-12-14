package session_redis

import (
	"context"
	"cs/class/session"
	"cs/meta"
	"encoding/json"
	"fmt"
	"framework/class/stream"
	"github.com/go-redis/redis/v8"
)

func keyOfSessionInfo(sessionId string) string {
	return fmt.Sprintf("cs:session:info:%s", sessionId)
}
func keyOfSessionClients(sessionId string) string {
	return fmt.Sprintf("cs:session:clients:%s", sessionId)
}
func keyOfSessionRecords(sessionId string) string {
	return fmt.Sprintf("cs:sessin:records:%s", sessionId)
}

var _ session.Session = &plugin{}

type plugin struct {
	redisClient  *redis.Client
	streamClient stream.Stream
	opts         *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) SetEnable(sessionId string, enable bool) (bool, error) {
	key := keyOfSessionInfo(sessionId)

	// 查询会话ID是否存在
	{
		exists, err := p.ExistsInfo(sessionId)
		if err != nil {
			return false, err
		}
		if !exists {
			return false, fmt.Errorf("会话信息不存在")
		}
	}

	// 设置Enable
	{
		_, err := p.redisClient.HSet(context.TODO(), key, "Enable", enable).Result()
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (p *plugin) GetEnable(sessionId string) (bool, error) {
	key := keyOfSessionInfo(sessionId)

	{
		res, err := p.redisClient.HGet(context.TODO(), key, "Enable").Result()
		if err != nil {
			return false, err
		}
		if res == "0" {
			return false, nil
		}
		if res == "1" {
			return true, nil
		}
		return false, fmt.Errorf("会话信息Enable值错误(%s)", res)
	}
}

func (p *plugin) SetInfo(session meta.Session) error {
	key := keyOfSessionInfo(session.GetSessionId())

	{
		j, err := json.Marshal(session)
		if err != nil {
			return err
		}
		_, err = p.redisClient.Set(context.TODO(), key, j, 0).Result()
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *plugin) GetInfo(sessionId string, session meta.Session) (bool, error) {
	key := keyOfSessionInfo(sessionId)

	{
		res, err := p.redisClient.Get(context.TODO(), key).Result()
		if err != nil {
			return false, err
		}
		if err := json.Unmarshal([]byte(res), session); err != nil {
			return false, err
		}
	}

	return true, nil
}

func (p *plugin) ExistsInfo(sessionId string) (bool, error) {
	key := keyOfSessionInfo(sessionId)
	num, err := p.redisClient.Exists(context.TODO(), key).Result()
	if err != nil {
		return false, err
	}
	if num == 0 {
		return false, nil
	}
	if num == 1 {
		return true, nil
	}
	return false, fmt.Errorf("会话信息num(%d)超出判断值", num)
}

func (p *plugin) SetClient(session meta.Session, client meta.Client) error {
	key := keyOfSessionClients(session.GetSessionId())

	{
		j, err := json.Marshal(client)
		if err != nil {
			return err
		}
		_, err = p.redisClient.HSet(context.TODO(), key, client.GetUUID(), j).Result()
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *plugin) DelClient(session meta.Session, client meta.Client) error {
	key := keyOfSessionClients(session.GetSessionId())

	{
		_, err := p.redisClient.HDel(context.TODO(), key, client.GetUUID()).Result()
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *plugin) ExistsClient(session meta.Session, client meta.Client) (bool, error) {
	key := keyOfSessionInfo(session.GetSessionId())

	{
		exists, err := p.redisClient.HExists(context.TODO(), key, client.GetUUID()).Result()
		if err != nil {
			return false, err
		}
		return exists, nil
	}
}

func (p *plugin) GetAllClients(session meta.Session, clients []interface{}) error {
	panic("implement me")
}

func (p *plugin) PushMessage(session meta.Session, message []byte) (string, error) {
	topic := keyOfSessionRecords(session.GetSessionId())
	return p.streamClient.Push(topic, message)
}

func (p *plugin) PullMessage(session meta.Session, lastMessageId string, count uint64) ([][]byte, error) {
	topic := keyOfSessionRecords(session.GetSessionId())
	res, err := p.streamClient.Pull(topic, lastMessageId, count)
	if err != nil {
		return nil, err
	}
	rows := make([][]byte, 0)
	for _, row := range res {
		rows = append(rows, row.Message())
	}
	return rows, nil
}
