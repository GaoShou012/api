package global

import (
	"framework/class/logger"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var (
	RedisClient *redis.Client
	DBMaster    *gorm.DB
	DBSlave     *gorm.DB
	Logger      logger.Logger
)
