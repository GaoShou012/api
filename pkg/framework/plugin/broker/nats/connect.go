package broker_nats

import (
	"framework/utils"
	"github.com/go-redis/redis"
)

func connect(dns string) (*redis.Client, error) {
	return utils.RedisClient(dns)
}
