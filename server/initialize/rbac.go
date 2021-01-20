package initialize

import (
	"api/global"
	"api/models"
	"framework/class/rbac"
	"framework/plugin/rbac/rbac_mysql_redis"
	"framework/plugin/rbac/rbac_mysql_redis/api_adapter"
	"framework/plugin/rbac/rbac_mysql_redis/menu_adapter"
	"framework/plugin/rbac/rbac_mysql_redis/role_adapter"
)

func InitRBAC() {
	var roleAdapter rbac.RoleAdapter
	var menuAdapter rbac.MenuAdapter
	var apiAdapter rbac.ApiAdapter

	{
		callback := &api_adapter.Callback{
			Authority: nil,
			SelectByMethodAndPathForAuthority: func(operator rbac.Operator, method string, path string) (rbac.Api, error) {
				model := &models.RbacApi{
					Model:  models.Model{},
					Method: nil,
					Path:   nil,
				}
				res := global.DBSlave.Table(model.GetTableName()).Where("method = ? and path = ?", method, path).Find(model)
				if res.Error != nil {
					return nil, res.Error
				}
				return model, nil
			},
		}
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
			api_adapter.WithModel(&models.RbacApi{}),
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
			GetMenuGroupIdByRoleIdMulti: func(operator rbac.Operator, roleIdMulti []uint64) ([]uint64, error) {
				rows := make([]models.RbacRoleAssocMenuGroup, 0)
				res := global.DBSlave.Table((&models.RbacRoleAssocMenuGroup{}).GetTableName()).Select("menu_group_id").Where("role_id in (?)", roleIdMulti).Find(&rows)
				if res.Error != nil {
					return nil, res.Error
				}
				ids := make([]uint64, len(rows))
				for index, row := range rows {
					ids[index] = *row.MenuGroupId
				}
				return ids, nil
			},
			GetMenuIdByRoleIdMulti: func(operator rbac.Operator, roleIdMulti []uint64) ([]uint64, error) {
				rows := make([]models.RbacRoleAssocMenu, 0)
				res := global.DBSlave.Table((&models.RbacRoleAssocMenu{}).GetTableName()).Select("menu_id").Where("role_id in (?)", roleIdMulti).Find(&rows)
				if res.Error != nil {
					return nil, res.Error
				}
				ids := make([]uint64, len(rows))
				for index, row := range rows {
					ids[index] = *row.MenuId
				}
				return ids, nil
			},
		}
		menuAdapter = menu_adapter.New(
			menu_adapter.WithModel(&models.RbacMenu{}, &models.RbacMenuGroup{}),
			menu_adapter.WithCallback(callback),
			menu_adapter.WithGorm(global.DBMaster, global.DBSlave),
		)
	}

	{
		callback := &role_adapter.Callback{
			Authority: role_adapter.Authority(func(operator rbac.Operator, roleId uint64) (bool, error) {
				return true, nil
			}),
			IsRoleAssocApiById: func(roleId uint64, apiId uint64) bool {
				count := 0
				res := global.DBSlave.Table((&models.RbacRoleAssocApi{}).GetTableName()).Where("role_id=? and api_id=?", roleId, apiId).Count(&count)
				if res.Error != nil {
					global.Logger.Error(res.Error)
					return false
				}
				if count == 0 {
					return false
				}
				return true
			},
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
				_role := role.(*models.RbacRole)
				_api := api.(*models.RbacApi)
				return &models.RbacRoleAssocApi{
					Model: models.Model{},
					//TenantId: _role.TenantId,
					RoleId: _role.Id,
					ApiId:  _api.Id,
				}
			}),
			AssocMenuGroup: role_adapter.AssocMenuGroup(func(role rbac.Role, group rbac.MenuGroup) rbac.Model {
				_role := role.(*models.RbacRole)
				_menuGroup := group.(*models.RbacMenuGroup)
				return &models.RbacRoleAssocMenuGroup{
					Model:       models.Model{},
					RoleId:      _role.Id,
					MenuGroupId: _menuGroup.Id,
				}
			}),
			AssocMenu: role_adapter.AssocMenu(func(role rbac.Role, menu rbac.Menu) rbac.Model {
				_role := role.(*models.RbacRole)
				_menu := menu.(*models.RbacMenu)
				return &models.RbacRoleAssocMenu{
					Model:  models.Model{},
					RoleId: _role.Id,
					MenuId: _menu.Id,
				}
			}),
		}
		roleAdapter = role_adapter.New(
			role_adapter.WithCallback(callback),
			role_adapter.WithGorm(global.DBMaster, global.DBSlave),
			role_adapter.WithModel(&models.RbacRole{}, &models.RbacRoleAssocApi{}, &models.RbacMenuGroup{}, &models.RbacMenu{}),
		)
	}

	global.RBAC = rbac_mysql_redis.New(
		rbac_mysql_redis.WithAdapter(roleAdapter, apiAdapter, menuAdapter),
	)
}
