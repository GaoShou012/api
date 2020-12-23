package api_adapter

import "framework/class/rbac"

type Callback struct {
	Authority
}

type Authority func(operator rbac.Operator, apiId uint64) (bool, error)
