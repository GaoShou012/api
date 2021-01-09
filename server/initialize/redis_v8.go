package initialize

import (
	"api/config"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

func ConnectRedisV8(conf *config.Redis) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               conf.Addr,
		Dialer:             nil,
		OnConnect:          nil,
		Username:           "",
		Password:           conf.Password,
		DB:                 0,
		MaxRetries:         2,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        time.Second * 10,
		ReadTimeout:        time.Second * 30,
		WriteTimeout:       time.Second * 30,
		PoolSize:           int(conf.PoolMax),
		MinIdleConns:       int(conf.PoolMin),
		MaxConnAge:         0,
		PoolTimeout:        time.Second * 30,
		IdleTimeout:        time.Minute * 5,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
		Limiter:            nil,
	})

	if _, err := redisClient.Ping(context.TODO()).Result(); err != nil {
		return nil, err
	}

	return redisClient, nil
}