package menu_adapter

import "framework/class/rbac"

type Callback struct {
	AuthorityMenuId
	AuthorityMenuGroupId
}

// 校验操作者 是否有权限操作 菜单ID
type AuthorityMenuId func(operator rbac.Operator, menuId uint64) (bool, error)

// 校验操作者 是否有权限操作 菜单组ID
type AuthorityMenuGroupId func(operator rbac.Operator, menuGroupId uint64) (bool, error)
