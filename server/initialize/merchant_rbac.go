package initialize

import (
	"api/global"
	"api/models"
	models_tenant "api/models/merchant"
	"framework/class/rbac"
	"framework/plugin/rbac/rbac_mysql_redis"
	"framework/plugin/rbac/rbac_mysql_redis/api_adapter"
	"framework/plugin/rbac/rbac_mysql_redis/menu_adapter"
	"framework/plugin/rbac/rbac_mysql_redis/role_adapter"
)

func InitTenantRBAC() {
	var roleAdapter rbac.RoleAdapter
	var menuAdapter rbac.MenuAdapter
	var apiAdapter rbac.ApiAdapter

	global.TenantRBAC = rbac_mysql_redis.New(
		rbac_mysql_redis.WithAdapter(roleAdapter, apiAdapter, menuAdapter),
	)

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
		apiAdapter = api_adapter.New(
			api_adapter.WithModel(&models_tenant.RbacApi{}),
			api_adapter.WithGorm(global.DBMaster, global.DBSlave),
			api_adapter.WithRedisClient(global.RedisClient),
			api_adapter.WithCallback(callback),
		)
	}

	{
		callback := &menu_adapter.Callback{
			AuthorityMenuId: menu_adapter.AuthorityMenuId(func(operator rbac.Operator, menuId uint64) (bool, error) {
				return true, nil
			}),
			AuthorityMenuGroupId: menu_adapter.AuthorityMenuGroupId(func(operator rbac.Operator, menuGroupId uint64) (bool, error) {
				return true, nil
			}),
		}
		menuAdapter = menu_adapter.New(
			menu_adapter.WithModel(&models_tenant.RbacMenu{}, &models_tenant.RbacMenuGroup{}),
			menu_adapter.WithCallback(callback),
			menu_adapter.WithGorm(global.DBMaster, global.DBSlave),
		)
	}

	{
		callback := &role_adapter.Callback{
			Authority: role_adapter.Authority(func(operator rbac.Operator, roleId uint64) (bool, error) {
				return true, nil
			}),
			ExistsByRoleIdReqMethodAndPath: role_adapter.ExistsByRoleIdReqMethodAndPath(func(roleId uint64, method string, path string) (bool, error) {
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
			}),
			AssocApi: role_adapter.AssocApi(func(role rbac.Role, api rbac.Api) rbac.Model {
				_role := role.(*models_tenant.RbacRole)
				_api := api.(*models_tenant.RbacApi)
				return &models_tenant.RbacRoleAssocApi{
					RbacRoleAssocApi: models.RbacRoleAssocApi{
						Model:  models.Model{},
						RoleId: _role.Id,
						ApiId:  _api.Id,
					},
				}
			}),
			AssocMenuGroup: role_adapter.AssocMenuGroup(func(role rbac.Role, group rbac.MenuGroup) rbac.Model {
				_role := role.(*models_tenant.RbacRole)
				_menuGroup := group.(*models_tenant.RbacMenuGroup)
				return &models_tenant.RbacRoleAssocMenuGroup{
					RbacRoleAssocMenuGroup: models.RbacRoleAssocMenuGroup{
						Model:       models.Model{},
						RoleId:      _role.Id,
						MenuGroupId: _menuGroup.Id,
					},
				}
			}),
			AssocMenu: role_adapter.AssocMenu(func(role rbac.Role, menu rbac.Menu) rbac.Model {
				_role := role.(*models_tenant.RbacRole)
				_menu := menu.(*models_tenant.RbacMenu)
				return &models_tenant.RbacRoleAssocMenu{
					RbacRoleAssocMenu: models.RbacRoleAssocMenu{
						Model:  models.Model{},
						RoleId: _role.Id,
						MenuId: _menu.Id,
					},
				}
			}),
		}
		roleAdapter = role_adapter.New(
			role_adapter.WithCallback(callback),
			role_adapter.WithGorm(global.DBMaster, global.DBSlave),
			role_adapter.WithModel(&models_tenant.RbacRole{}, &models_tenant.RbacRoleAssocApi{}, &models_tenant.RbacMenuGroup{}, &models_tenant.RbacMenu{}),
		)
	}
}
