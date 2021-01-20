package utils

import (
	"github.com/go-redis/redis/v7"
	"net/url"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func RedisClusterClientV7(dns string) (*redis.ClusterClient, error) {
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

	cli := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              strings.Split(addr, ","),
		MaxRedirects:       0,
		ReadOnly:           false,
		RouteByLatency:     false,
		RouteRandomly:      false,
		ClusterSlots:       nil,
		OnNewNode:          nil,
		OnConnect:          nil,
		Password:           password,
		MaxRetries:         2,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        30 * time.Second,
		ReadTimeout:        3 * time.Second,
		WriteTimeout:       30 * time.Second,
		PoolSize:           poolMax,
		MinIdleConns:       poolMin,
		MaxConnAge:         0,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        5 * time.Minute,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})
	return cli, nil
}
