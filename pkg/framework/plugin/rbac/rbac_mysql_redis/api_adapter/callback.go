package api_adapter

import "framework/class/rbac"

type Callback struct {
	Authority
}

/*
	鉴权
	校验操作者，是否有权限操作此API
*/
type Authority func(operator rbac.Operator, apiId uint64) (bool, error)
