package initialize

import (
	"api/config"
	"api/global"
	"framework/utils"
	"github.com/go-redis/redis"
)

func InitRedis(conf *config.Redis) error {
	client, err := ConnectRedis(conf)
	global.RedisClient = client
	return err
}

func ConnectRedis(conf *config.Redis) (*redis.Client, error) {
	redisClient, err := utils.RedisClient(conf.DNS)
	if err != nil {
		return nil, err
	}

	if _, err := redisClient.Ping().Result(); err != nil {
		return nil, err
	}

	return redisClient, nil
	//redisClient := redis.NewClient(&redis.Options{
	//	Network:            "",
	//	Addr:               conf.Addr,
	//	Dialer:             nil,
	//	OnConnect:          nil,
	//	Password:           conf.Password,
	//	DB:                 0,
	//	MaxRetries:         2,
	//	MinRetryBackoff:    0,
	//	MaxRetryBackoff:    0,
	//	DialTimeout:        time.Second * 10,
	//	ReadTimeout:        time.Second * 30,
	//	WriteTimeout:       time.Second * 30,
	//	PoolSize:           int(conf.PoolMax),
	//	MinIdleConns:       int(conf.PoolMin),
	//	MaxConnAge:         0,
	//	PoolTimeout:        time.Second * 30,
	//	IdleTimeout:        time.Minute * 5,
	//	IdleCheckFrequency: 0,
	//	TLSConfig:          nil,
	//})
	//
	//if _, err := redisClient.Ping().Result(); err != nil {
	//	return nil, err
	//}
	//
	//return redisClient, nil
}
