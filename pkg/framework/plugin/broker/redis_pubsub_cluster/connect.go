package broker_redis_pubsub_cluster

import (
	"framework/utils"
	"github.com/go-redis/redis"
)

func connect(dns string) (*redis.ClusterClient, error) {
	return utils.RedisClusterClient(dns)
}
