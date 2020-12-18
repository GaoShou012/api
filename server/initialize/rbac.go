package initialize

import (
	"api/global"
	"framework/env"
	"framework/plugin/rbac/rbac_mysql_redis"
	"github.com/jinzhu/gorm"
)

func InitRBAC(db *gorm.DB) {
	env.ApiAdapter = rbac_mysql_redis.NewApiAdapter(
		rbac_mysql_redis.WithDatabase(db),
	)
	global.RBAC.ApiAdapter = env.ApiAdapter
}
