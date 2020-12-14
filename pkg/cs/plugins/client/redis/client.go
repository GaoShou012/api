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

var _ client.Client = &Client{}

type Client struct {
	redisClient *redis.Client
}

func (c *Client) Init() error {
	return nil
}

func (c *Client) SetInfo(uuid string, client meta.Client) error {
	key := keyOfClientInfo(uuid)

	{
		j, err := json.Marshal(client)
		if err != nil {
			return err
		}
		_, err = c.redisClient.Set(context.TODO(), key, j, 0).Result()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) GetInfo(uuid string, client meta.Client) (bool, error) {
	key := keyOfClientInfo(uuid)

	{
		res, err := c.redisClient.Get(context.TODO(), key).Result()
		if err != nil {
			return false, err
		}
		if err := json.Unmarshal([]byte(res), client); err != nil {
			return false, err
		}
	}

	return true, nil
}

func (c *Client) ExistsInfo(uuid string) (bool, error) {
	key := keyOfClientInfo(uuid)
	num, err := c.redisClient.Exists(context.TODO(), key).Result()
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

func (c *Client) SetSession(client meta.Client, sessionId string, session meta.Session) error {
	key := keyOfClientSessions(client.GetUUID())

	{
		j, err := json.Marshal(session)
		if err != nil {
			return err
		}
		_, err = c.redisClient.HSet(context.TODO(), key, sessionId, j).Result()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) DelSession(client meta.Client, sessionId string) error {
	key := keyOfClientSessions(client.GetUUID())

	{
		_, err := c.redisClient.HDel(context.TODO(), key, sessionId).Result()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) ExistsSession(client meta.Client, sessionId string) (bool, error) {
	key := keyOfClientSessions(client.GetUUID())

	{
		exists, err := c.redisClient.HExists(context.TODO(), key, sessionId).Result()
		if err != nil {
			return false, err
		}
		return exists, nil
	}
}

func (c *Client) GetAllSessions(client meta.Client, sessions interface{}) error {
	panic("implement me")
}
