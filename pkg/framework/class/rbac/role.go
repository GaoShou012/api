package rbac

type Role interface{}

type RoleAdapter interface {
	CreateRole(operator Operator, role Role) error
	DeleteRole(operator Operator, roleId uint64) error
	UpdateRole(operator Operator, role Role) error
	SelectById(roleId uint64) (Role, error)
	IncrReference(operator Operator, roleId uint64) error

	/*
		校验操作者，是否有权限，操作此角色ID
	*/
	VerifyIdWithOperator(roleId uint64,operator Operator) (bool,error)

	/*
		角色关联菜单组
		增加菜单组的饮用量
	*/
	AssocMenuGroup(role Role, group MenuGroup) error

	/*
		角色关联API
	*/
	AssocApi(role Role,api Api) error
}
