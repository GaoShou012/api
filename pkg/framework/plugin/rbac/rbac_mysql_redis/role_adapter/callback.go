package role_adapter

import "framework/class/rbac"

type Callback struct {
	Authority
	ExistsByRoleIdReqMethodAndPath
}

type Authority func(operator rbac.Operator, roleId uint64) (bool, error)
type ExistsByRoleIdReqMethodAndPath func(roleId uint64,method string,path string) (bool,error)