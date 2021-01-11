package broker_redis_pubsub_v8

import (
	"framework/utils"
	"github.com/go-redis/redis/v8"
)

func connect(dns string) (*redis.Client, error) {
	return utils.RedisClientV8(dns)
}
