package client

import (
	"context"
	"cs/class/client"
	"cs/meta"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func keyOfSessionsOfClient(uuid string) string {
	return fmt.Sprintf("cs:client:sessions:%s", uuid)
}

var _ client.Client = &Client{}

type Client struct {
	redisClient *redis.Client
}

func (c *Client) AddSession(clientInfo meta.Client, session meta.Session) (bool, error) {
	key := keyOfSessionsOfClient(clientInfo.GetUUID())
	res, err := c.redisClient.HSet(context.TODO(), key, session.GetSessionId()).Result()
	if err != nil {
		return false, err
	}

}

func (c *Client) DelSession(clientInfo meta.Client, sessionId string) (bool, error) {
	key := keyOfSessionsOfClient(clientInfo.GetUUID())
	res, err := c.redisClient.HDel(context.TODO(), key, sessionId).Result()

}

func (c *Client) ExistsSession(clientInfo client.Info, sessionId string) (bool, error) {
	panic("implement me")
}
