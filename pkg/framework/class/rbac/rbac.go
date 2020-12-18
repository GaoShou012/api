package rbac

type RBAC interface {
	CreateRole(role Role) error
	UpdateRole(role Role) error
	DeleteRole(roleId uint64) error

	EnforcerRole()

	/*
		权限判断
	*/
	Enforcer(authorityId string, method string, path string) (bool, error)
}