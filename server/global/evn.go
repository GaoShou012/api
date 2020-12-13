package global

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

var (
	CasbinAdapter *gormadapter.Adapter
	RedisClient   *redis.Client
	//RedisClient *redis.ClusterClient
	DBMaster *gorm.DB
	DBSlave  *gorm.DB
)
