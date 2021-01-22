package stream_redis_stream

import (
	"framework/utils"
	"github.com/go-redis/redis/v7"
)

func connect(dns string) (*redis.ClusterClient, error) {
	return utils.RedisClusterClientV7(dns)
}
