package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

var (
	CasbinEnforcer *casbin.Enforcer
	RedisClient    *redis.Client
	//RedisClient *redis.ClusterClient
	DBMaster *gorm.DB
	DBSlave  *gorm.DB
)
