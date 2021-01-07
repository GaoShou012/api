package redis_sortdset

import "github.com/go-redis/redis"

type Options struct {
	redisClient *redis.Client
	//Z *redis_sortdset.Z
}
