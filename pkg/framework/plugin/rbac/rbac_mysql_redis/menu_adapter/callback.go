package api_adapter

import "framework/class/rbac"

type Callback struct {
	AuthorityMenuId
	AuthorityMenuGroupId
}

type AuthorityMenuId func(operator rbac.Operator, menuId uint64) (bool, error)
type AuthorityMenuGroupId func(operator rbac.Operator, menuGroupId uint64) (bool, error)
