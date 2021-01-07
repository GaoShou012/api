package lock_redis_set

import (
	"context"
	"fmt"
	"framework/class/lock"
	"time"
)

var _ lock.Lock = &plugin{}

type plugin struct {
	key         string
	val         string
	opts        *Options
}

func (p *plugin) Init() error {
	if p.opts.redisClient == nil {
		return fmt.Errorf("Must init redis_sortdset client.\n")
	}
	return nil
}

func (p *plugin) Lock(key string, val string, timeout time.Duration) (bool, error) {
	return p.opts.redisClient.SetNX(context.TODO(), key, val, timeout).Result()
}
func (p *plugin) Unlock(key string, val string) (bool, error) {
	_, err := p.opts.redisClient.Del(context.TODO(), key).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}
