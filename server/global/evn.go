package global

import (
	"framework/class/logger"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

var (
	RedisClient *redis.Client
	DBMaster    *gorm.DB
	DBSlave     *gorm.DB
	Logger      logger.Logger
)
