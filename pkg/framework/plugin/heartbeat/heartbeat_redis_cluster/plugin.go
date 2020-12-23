package heartbeat_redis

import (
	"context"
	"errors"
	"fmt"
	"framework/class/heartbeat"
	"github.com/go-redis/redis/v8"
	"time"
)

var _ heartbeat.Heartbeat = &plugin{}

type plugin struct {
	redisClient *redis.Client
	timeout     time.Duration
	opts        *Options
}

func (p *plugin) key(uuid string) string {
	return fmt.Sprintf("heartbeat:%s", uuid)
}

func (p *plugin) Init() error {
	// 必须要配置redis
	if p.opts.redisClient == nil {
		return errors.New("redis can't be null\n")
	}
	p.redisClient = p.opts.redisClient

	// 默认超时时间10秒
	if p.opts.timeout == 0 {
		p.opts.timeout = time.Second * 10
	}
	p.timeout = p.opts.timeout

	return nil
}

func (p *plugin) Heartbeat(uuid string) error {
	_, err := p.redisClient.Set(context.TODO(), p.key(uuid), time.Now().String(), p.timeout).Result()
	return err
}

func (p *plugin) IsAlive(uuid string) (bool, error) {
	num, err := p.redisClient.Exists(context.TODO(), p.key(uuid)).Result()
	if err != nil {
		return false, err
	}
	if num == 0 {
		return false, nil
	}
	if num == 1 {
		return true, nil
	}
	return false, fmt.Errorf("存在的key(%s)数量错误:%d", p.key(uuid), num)
}
