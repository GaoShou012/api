package initialize

import (
	"api/global"
	"api/models"
	"framework/class/rbac"
	"framework/env"
	"framework/plugin/rbac/rbac_mysql_redis/api_adapter"
	"framework/plugin/rbac/rbac_mysql_redis/role_adapter"
	"github.com/jinzhu/gorm"
)

func InitRBAC(db *gorm.DB) {
	{
		callback := &api_adapter.Callback{}
		callback.Authority = func(operator rbac.Operator, apiId uint64) (bool, error) {
			//tmp,err := env.ApiAdapter.SelectById(apiId)
			//		//if err != nil {
			//		//	return false,err
			//		//}
			//		//api :=  tmp.(*models.AuthorityApis)
			//		//
			//		//if operator.GetTenantId() != *api.TenantId {
			//		//	return false,nil
			//		//}
			return true, nil
		}
		env.ApiAdapter = api_adapter.New(
			api_adapter.WithModel(&models.AuthorityApis{}),
			api_adapter.WithGorm(global.DBMaster, global.DBSlave),
			api_adapter.WithRedisClient(global.RedisClient),
			api_adapter.WithCallback(callback),
		)
	}

	{
		callback := &role_adapter.Callback{}
		callback.Authority = func(operator rbac.Operator, roleId uint64) (bool, error) {
			return true, nil
		}
		callback.ExistsByRoleIdReqMethodAndPath = func(roleId uint64, method string, path string) (bool, error) {
			count := 0
			res := global.DBSlave.Where("role_id = ? and method =? and path =?", roleId, method, path).Count(&count)
			if res.Error != nil {
				return false, res.Error
			}
			if count == 0 {
				return false, nil
			} else {
				return true, nil
			}
		}
		env.RoleAdapter = role_adapter.New(
			role_adapter.WithCallback(callback),
		)
	}
}
