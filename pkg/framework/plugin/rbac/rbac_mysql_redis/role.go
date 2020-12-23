package rbac_mysql_redis

import "framework/class/rbac"

var _ rbac.RoleAdapter = &RoleAdapter{}

type RoleAdapter struct {
}

func (r RoleAdapter) CreateRole(operator rbac.Operator, role rbac.Role) error {
	panic("implement me")
}

func (r RoleAdapter) DeleteRole(operator rbac.Operator, roleId uint64) error {
	panic("implement me")
}

func (r RoleAdapter) UpdateRole(operator rbac.Operator, role rbac.Role) error {
	panic("implement me")
}

func (r RoleAdapter) SelectById(roleId uint64) (rbac.Role, error) {
	panic("implement me")
}

func (r RoleAdapter) IncrReference(operator rbac.Operator, roleId uint64) error {
	panic("implement me")
}

func (r RoleAdapter) VerifyIdWithOperator(roleId uint64, operator rbac.Operator) (bool, error) {
	panic("implement me")
}

func (r RoleAdapter) AssocMenuGroup(role rbac.Role, group rbac.MenuGroup) error {
	panic("implement me")
}

func (r RoleAdapter) AssocApi(role rbac.Role, api rbac.Api) error {
	panic("implement me")
}

