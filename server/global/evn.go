package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

var (
	RedisClient        *redis.Client
	//RedisClusterClient *redis.ClusterClient
	DBMaster           *gorm.DB
	DBSlave            *gorm.DB
)
