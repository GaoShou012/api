package utils

import (
	"github.com/go-redis/redis"
	"net/url"
	"runtime"
	"strconv"
	"time"
)

func RedisClient(dns string) (*redis.Client, error) {
	u, err := url.Parse(dns)
	if err != nil {
		return nil, err
	}

	addr := u.Host
	//username := u.User.Username()
	password, _ := u.User.Password()

	params, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}

	var poolMax int
	{
		arr, ok := params["PoolMax"]
		if !ok {
			poolMax = runtime.NumCPU()
		} else {
			num, err := strconv.Atoi(arr[0])
			if err != nil {
				return nil, err
			}
			poolMax = num
		}
	}

	var poolMin int
	{
		arr, ok := params["PoolMin"]
		if !ok {
			poolMin = runtime.NumCPU()
		} else {
			num, err := strconv.Atoi(arr[0])
			if err != nil {
				return nil, err
			}
			poolMin = num
		}
	}

	// new redis_sortdset client
	redisClient := redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               addr,
		Dialer:             nil,
		OnConnect:          nil,
		Password:           password,
		DB:                 0,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        10 * time.Second,
		ReadTimeout:        30 * time.Second,
		WriteTimeout:       30 * time.Second,
		PoolSize:           poolMax,
		MinIdleConns:       poolMin,
		MaxConnAge:         0,
		PoolTimeout:        time.Second * 30,
		IdleTimeout:        time.Minute * 5,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})

	// ping redis_sortdset server
	if _, err := redisClient.Ping().Result(); err != nil {
		return nil, err
	}

	return redisClient, nil
}
