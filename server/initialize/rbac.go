package initialize

import (
	"github.com/jinzhu/gorm"
)

func InitRBAC(db *gorm.DB) {
	//env.ApiAdapter = rbac_mysql_redis.New(
	//	rbac_mysql_redis.WithDatabase(db),
	//)
	//global.RBAC.ApiAdapter = env.ApiAdapter
}
