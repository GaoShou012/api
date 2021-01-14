package rbac

type RBAC interface {
	/*
		创建API
	*/
	CreateApi(operator Operator, api Api) error

	/*
		删除API
	*/
	DeleteApi(operator Operator, apiId uint64) (bool, error)

	/*
		更新API
	*/
	UpdateApi(operator Operator, apiId uint64, api Api) error

	/*
		创建菜单
	*/
	CreateMenu(operator Operator, menuGroupId uint64, menu Menu) error

	/*
		删除菜单
	*/
	DeleteMenu(operator Operator, menuId uint64) (bool, error)

	/*
		更新菜单
	*/
	UpdateMenu(operator Operator, menuId uint64, menu Menu) error

	/*
		创建菜单组
	*/
	CreateMenuGroup(operator Operator, group MenuGroup) error

	/*
		删除菜单组
	*/
	DeleteMenuGroup(operator Operator, menuGroupId uint64) (bool, error)

	/*
		更新菜单组
	*/
	UpdateMenuGroup(operator Operator, menuGroupId uint64, group MenuGroup) error

	/*
		创建角色
	*/
	CreateRole(operator Operator, role Role) error

	/*
		删除角色
	*/
	DeleteRole(operator Operator, roleId uint64) (bool, error)

	/*
		更新角色
	*/
	UpdateRole(operator Operator, roleId uint64, role Role) error

	/*
		查询角色
	*/
	SelectRoles(operator Operator, rolesId string) ([]Role, error)

	/*
		角色关联API
	*/
	RoleAssocApi(operator Operator, roleId uint64, apiId uint64) error

	/*
		角色取消关联API
	*/
	RoleDisassociateApi(operator Operator, assocId uint64) error

	/*
		角色关联菜单
	*/
	RoleAssocMenu(operator Operator, roleId uint64, menuId uint64) error

	/*
		角色取消关联菜单
	*/
	RoleDisassociateMenu(operator Operator, assocId uint64) error

	/*
		角色关联菜单组
	*/
	RoleAssocMenuGroup(operator Operator, roleId uint64, menuGroupId uint64) error

	/*
		角色取消关联菜单组
	*/
	RoleDisassociateMenuGroup(operator Operator, assocId uint64) error

	/*
		权限判断
	*/
	Enforcer(authorityId string, method string, path string) (bool, error)
}
