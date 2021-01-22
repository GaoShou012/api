package api_adapter

import "framework/class/rbac"

type Callback struct {
	/*
		鉴权
		校验操作者，是否有权限操作此API
	*/
	Authority func(operator rbac.Operator, apiId uint64) (bool, error)

	// 根据方法和路径，查询API，用于权限校验
	SelectByMethodAndPathForAuthority func(operator rbac.Operator, method string, path string) (rbac.Api, error)

	ExistsByMethodAndPath func(operator rbac.Operator,method string,path string) (bool,error)
}
