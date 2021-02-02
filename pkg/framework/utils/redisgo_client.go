package utils

import (
	"github.com/garyburd/redigo/redis"
	"net/url"
	"runtime"
	"strconv"
	"time"
)

func RedisGoConnect(dns string) (*redis.Pool, error) {
	u, err := url.Parse(dns)
	if err != nil {
		return nil, err
	}

	addr := u.Host
	password, _ := u.User.Password()
	params, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
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

	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if password != "" {
				_, err := conn.Do("AUTH", password)
				if err != nil {
					conn.Close()
					return nil, err
				}
			}
			return conn, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         poolMin,
		MaxActive:       poolMax,
		IdleTimeout:     time.Minute * 5,
		Wait:            true,
		MaxConnLifetime: 0,
	}

	return redisPool, nil
}
