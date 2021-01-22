package menu_adapter

import "framework/class/rbac"

type Callback struct {
	AuthorityMenuId
	AuthorityMenuGroupId

	SelectMenuGroupWithFieldsByRoleIdMulti
	SelectMenuWithFieldsByMenuGroupIdMulti

	// 根据多个角色ID，获取关联的菜单组ID
	GetMenuGroupIdByRoleIdMulti func(operator rbac.Operator, roleIdMulti []uint64) ([]uint64, error)

	// 根据多个角色ID，获取关联的菜单ID
	GetMenuIdByRoleIdMulti func(operator rbac.Operator, roleIdMulti []uint64) ([]uint64, error)
}

// 校验操作者 是否有权限操作 菜单ID
type AuthorityMenuId func(operator rbac.Operator, menuId uint64) (bool, error)

// 校验操作者 是否有权限操作 菜单组ID
type AuthorityMenuGroupId func(operator rbac.Operator, menuGroupId uint64) (bool, error)

// 根据角色ID（多项）查询菜单组
type SelectMenuGroupWithFieldsByRoleIdMulti func(operator rbac.Operator, roleIdMulti []uint64, fields string) ([]rbac.MenuGroup, error)

// 根据菜单组ID（多项）查询菜单
type SelectMenuWithFieldsByMenuGroupIdMulti func(operator rbac.Operator, menuGroupIdMulti []uint64, fields string) ([]rbac.Menu, error)

