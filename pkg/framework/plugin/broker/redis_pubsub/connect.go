package broker_redis_pubsub

import (
	"framework/utils"
	"github.com/go-redis/redis"
)

func connect(dns string) (*redis.Client, error) {
	return utils.RedisClient(dns)
}
