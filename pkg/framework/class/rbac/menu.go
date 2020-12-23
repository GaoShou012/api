package rbac

type Menu interface {
	Model
}
type MenuGroup interface {
	Model
}

type MenuAdapter interface {
	/*
		校验操作者，是否有权限操作菜单
	*/
	AuthorityMenu(operator Operator, menuId uint64) (bool, error)
	CreateMenuGroup(group MenuGroup) error
	DeleteMenuGroup(groupId uint64) (bool, error)
	UpdateMenuGroup(groupId uint64, group MenuGroup) error
	SelectMenuGroupById(menuGroupId uint64) (MenuGroup, error)

	/*
		校验操作者，是否有权限操作菜单组
	*/
	AuthorityMenuGroup(operator Operator, groupId uint64) (bool, error)

	CreateMenu(menu Menu) error
	DeleteMenu(menuId uint64) (bool,error)
	UpdateMenu(menuId uint64, menu Menu) error
	SelectMenuById(menuId uint64) (Menu, error)
	SelectMenuByGroupId(groupId uint64) ([]Menu, error)

	/*
		根据角色ID
		查询用户下的所有关联菜单
	*/
	SelectByRoleId(roleId uint64) error
}
