package adapter_redis

import (
	"cs/class"
	"cs/meta"
	"github.com/go-redis/redis/v8"
)

var _ class.Adapter = &Adapter{}

type Adapter struct {
	redisClient *redis.Client
	opts        *Options
}

func (a *Adapter) Init() error {
	return nil
}

func (a Adapter) ExistsSession(session meta.Session) (bool, error) {

	panic("implement me")
}

func (a Adapter) SaveSession(session meta.Session) error {
	panic("implement me")
}

func (a Adapter) ReadSession(sessionId string, session meta.Session) (bool, error) {
	panic("implement me")
}

func (a Adapter) ExistsClient(clientId string) (bool, error) {
	panic("implement me")
}

func (a Adapter) SaveClient(client meta.Client) error {
	panic("implement me")
}

func (a Adapter) ReadClient(clientId string, client meta.Client) (bool, error) {
	panic("implement me")
}
