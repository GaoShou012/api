package rbac

type Menu interface {
	GetGroupId() uint64
}
type MenuGroup interface {
	GetMenuName() string
	GetGroupName() string
	GetSort() uint64
}
type MenuAdapter interface {
	CreateMenuGroup(operator Operator, group MenuGroup) error
	UpdateMenuGroup(groupId uint64, group MenuGroup) error
	//DeleteMenuGroup(operator Operator, groupId uint64) error
	SelectMenuGroup(operator Operator) ([]MenuGroup, error)
	SelectAllMenuGroup(tenantId uint64) ([]MenuGroup, error)

	/*
		校验操作者，是否有权限操作菜单组
	*/
	VerifyGroupIdWithOperator(operator Operator, groupId uint64) (bool, error)

	SelectMenuGroupById(groupId uint64) (MenuGroup, error)

	CreateMenu(groupId uint64, menu Menu) error
	DeleteMenu(tenantId uint64, menuId uint64) error
	UpdateMenu(menuId uint64, menu Menu) error
	SelectByGroupId(tenantId uint64, groupId uint64) ([]Menu, error)
	SelectMenuById(menuId uint64) (Menu, error)
	SelectMenuByGroupId(groupId uint64) ([]Menu, error)

	/*
		根据用户ID
		查询用户下的所有关联菜单
	*/
	SelectByRoleId(tenantId uint64, roleId uint64) error
}
