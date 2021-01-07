package broker_redis_pubsub_cluster

import (
	"context"
	"github.com/go-redis/redis/v8"
	"net/url"
	"runtime"
	"strconv"
	"time"
)

func connect(dns string) (*redis.ClusterClient, error) {
	u, err := url.Parse(dns)
	if err != nil {
		return nil, err
	}

	addr := u.Host
	username := u.User.Username()
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
	redisClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              addr,
		NewClient:          nil,
		MaxRedirects:       0,
		ReadOnly:           false,
		RouteByLatency:     false,
		RouteRandomly:      false,
		ClusterSlots:       nil,
		Dialer:             nil,
		OnConnect:          nil,
		Username:           "",
		Password:           "",
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})

	// ping redis_sortdset server
	if _, err := redisClient.Ping(context.TODO()).Result(); err != nil {
		return nil, err
	}

	return redisClient, nil
}
