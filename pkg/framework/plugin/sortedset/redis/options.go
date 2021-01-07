package redis

import "github.com/go-redis/redis"

type Options struct {
	redisClient *redis.Client
	//Z *redis.Z
}
