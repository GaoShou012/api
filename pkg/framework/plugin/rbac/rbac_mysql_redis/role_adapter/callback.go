package role_adapter

import "framework/class/rbac"

type Callback struct {
	Authority
	ExistsByRoleIdReqMethodAndPath

	AssocApi
	AssocMenuGroup
	AssocMenu
}

type Authority func(operator rbac.Operator, roleId uint64) (bool, error)
type ExistsByRoleIdReqMethodAndPath func(roleId uint64, method string, path string) (bool, error)

// 返回角色与API的关联模型
type AssocApi func(role rbac.Role, api rbac.Api) rbac.Model

// 返回角色与慈丹的关联模型
type AssocMenu func(role rbac.Role, menu rbac.Menu) rbac.Model

// 返回角色与菜单组的关联模型
type AssocMenuGroup func(role rbac.Role, group rbac.MenuGroup) rbac.Model
