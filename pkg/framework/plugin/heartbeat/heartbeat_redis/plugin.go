package heartbeat_redis

import (
	"errors"
	"fmt"
	"framework/class/heartbeat"
	"time"
)

var _ heartbeat.Heartbeat = &plugin{}

type plugin struct {
	opts *Options
}

func (p *plugin) key(uuid string) string {
	return fmt.Sprintf("heartbeat:%s", uuid)
}

func (p *plugin) Init() error {
	// 必须要配置redis
	if p.opts.redisClient == nil {
		return errors.New("redis_sortdset can't be null\n")
	}

	// 默认超时时间10秒
	if p.opts.timeout == 0 {
		p.opts.timeout = time.Second * 10
	}

	return nil
}

func (p *plugin) Heartbeat(uuid string) error {
	_, err := p.opts.redisClient.Set(p.key(uuid), time.Now().String(), p.opts.timeout).Result()
	return err
}

func (p *plugin) IsAlive(uuid string) (bool, error) {
	num, err := p.opts.redisClient.Exists(p.key(uuid)).Result()
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
