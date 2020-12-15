package client_redis

import (
	"context"
	"cs/class/client"
	"cs/meta"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func keyOfClientInfo(uuid string) string {
	return fmt.Sprintf("cs:client:info:%s", uuid)
}
func keyOfClientSessions(uuid string) string {
	return fmt.Sprintf("cs:client:sessions:%s", uuid)
}

var _ client.Client = &plugin{}

type plugin struct {
	redisClient *redis.Client
	opts        *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) SetInfo(uuid string, client meta.Client) error {
	key := keyOfClientInfo(uuid)

	{
		j, err := json.Marshal(client)
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

func (p *plugin) GetInfo(uuid string, client meta.Client) (bool, error) {
	key := keyOfClientInfo(uuid)

	{
		res, err := p.redisClient.Get(context.TODO(), key).Result()
		if err != nil {
			return false, err
		}
		if err := json.Unmarshal([]byte(res), client); err != nil {
			return false, err
		}
	}

	return true, nil
}

func (p *plugin) ExistsInfo(uuid string) (bool, error) {
	key := keyOfClientInfo(uuid)
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

func (p *plugin) SetSession(client meta.Client, sessionId string, session meta.Session) error {
	key := keyOfClientSessions(client.GetUUID())

	{
		j, err := json.Marshal(session)
		if err != nil {
			return err
		}
		_, err = p.redisClient.HSet(context.TODO(), key, sessionId, j).Result()
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *plugin) DelSession(client meta.Client, sessionId string) error {
	key := keyOfClientSessions(client.GetUUID())

	{
		_, err := p.redisClient.HDel(context.TODO(), key, sessionId).Result()
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *plugin) ExistsSession(client meta.Client, sessionId string) (bool, error) {
	key := keyOfClientSessions(client.GetUUID())

	{
		exists, err := p.redisClient.HExists(context.TODO(), key, sessionId).Result()
		if err != nil {
			return false, err
		}
		return exists, nil
	}
}

func (p *plugin) GetAllSessions(client meta.Client, sessions interface{}) error {
	panic("implement me")
}
