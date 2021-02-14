package id_generator_redis_v7

import (
	"fmt"
	"framework/class/id_generator"
	"strconv"
)

var _ id_generator.IdGenerator = &plugin{}

type plugin struct {
	opts *Options
}

func (p *plugin) Init() error {
	if p.opts.redisClient == nil {
		return fmt.Errorf("redis client can't is nil\n")
	}
	return nil
}

func (p *plugin) Incr(key string) (int64, error) {
	return p.opts.redisClient.Incr(key).Result()
}

func (p *plugin) Get(key string) (int64, error) {
	val, err := p.opts.redisClient.Get(key).Result()
	if err != nil {
		return 0, err
	}
	// 如果val是空字符串，返回默认值0
	if val == "" {
		return 0, nil
	}
	// 字符串转数字
	num, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return int64(num), nil
}

func (p *plugin) HIncr(key string, field string) (int64, error) {
	return p.opts.redisClient.HIncrBy(key, field, 1).Result()
}

func (p *plugin) HGet(key string, field string) (int64, error) {
	val, err := p.opts.redisClient.HGet(key, field).Result()
	if err != nil {
		return 0, err
	}
	// 如果val是空字符串，返回默认值0
	if val == "" {
		return 0, nil
	}
	num, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return int64(num), nil
}

func (p *plugin) HGetAll(key string) (map[string]int64, error) {
	res, err := p.opts.redisClient.HGetAll(key).Result()
	if err != nil {
		return nil, err
	}
	num := make(map[string]int64)
	for k, v := range res {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("转换数据类型失败:v=%s", v)
		}
		num[k] = int64(n)
	}
	return num, nil
}
