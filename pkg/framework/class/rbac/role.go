package rbac

type Role interface{}

type RoleAssocApi interface {
	GetRoleId() uint64
	GetApiId() uint64
}
type RoleAssocMenuGroup interface {
	GetRoleId() uint64
	GetMenuGroupId() uint64
}
type RoleAssocMenu interface {
	GetRoleId() uint64
	GetMenuId() uint64
}

type RoleAdapter interface {
	/*
		校验操作者，是否有权限，操作此角色ID
	*/
	Authority(operator Operator, roleId uint64) (bool, error)

	CreateRole(role Role) error
	DeleteRole(roleId uint64) (bool, error)
	UpdateRole(roleId uint64, role Role) error
	SelectById(roleId uint64) (Role, error)
	SelectByPersonId(personId uint64) ([]Role,error)

	/*
		根据ID，查询角色&API，关联数据
	*/
	SelectAssocApiById(assocId uint64) (RoleAssocApi, error)
	SelectAssocMenuById(assocId uint64) (RoleAssocMenu, error)
	SelectAssocMenuGroupById(assocId uint64) (RoleAssocMenuGroup, error)

	/*
		角色关联菜单组
		增加菜单组的饮用量
	*/
	AssocMenuGroup(role Role, group MenuGroup) error

	/*
		角色取消关联菜单组
	*/
	DisassociateMenuGroup(assocId uint64) (bool, error)

	/*
		角色关联菜单
	*/
	AssocMenu(role Role, menu Menu) error

	/*
		角色取消关联菜单
	*/
	DisassociateMenu(assocId uint64) (bool, error)

	/*
		角色关联API
	*/
	AssocApi(role Role, api Api) error

	/*
		角色取消关联API
	*/
	DisassociateApi(assocId uint64) (bool, error)

	/*
		角色是否存在API
	*/
	//EnforcerApi(roleId uint64, method string, path string) (bool, error)

	/*
		角色校验API权限
	*/
	EnforcerApi(roleId uint64,apiId uint64) bool
}
