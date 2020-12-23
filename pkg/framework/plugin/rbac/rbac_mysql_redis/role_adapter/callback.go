package api_adapter

import "framework/class/rbac"

type Callback struct {
	Authority
}

type Authority func(operator rbac.Operator, roleId uint64) (bool, error)
