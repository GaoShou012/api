package broker_redis_stream

import (
	"framework/utils"
	"github.com/go-redis/redis"
)

func connect(dns string) (*redis.Client, error) {
	return utils.RedisClient(dns)
}
