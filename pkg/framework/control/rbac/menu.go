package rbac

import (
	"errors"
	"framework/class/rbac"
	"framework/env"
)

/*
	创建菜单
*/
func CreateMenu(operator rbac.Operator, groupId uint64, menu rbac.Menu) error {
	// 校验操作者权限
	ok, err := env.MenuAdapter.VerifyGroupIdWithOperator(operator, groupId)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("权限不足")
	}

	// 创建菜单
	if err := env.MenuAdapter.CreateMenu(groupId, menu); err != nil {
		return err
	}

	return nil
}

/*
	更新菜单
*/
func UpdateMenu(operator rbac.Operator, menuId uint64, menu rbac.Menu) error {
	// 查询菜单组
	menu, err := env.MenuAdapter.SelectMenuById(menuId)
	if err != nil {
		return err
	}

	// 校验操作者权限
	ok, err := env.MenuAdapter.VerifyGroupIdWithOperator(operator, menu.GetGroupId())
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("权限不足")
	}

	// 更新菜单
	if err := env.MenuAdapter.UpdateMenu(menuId, menu); err != nil {
		return err
	}

	return nil
}

func SelectMenuByGroupId(operator rbac.Operator, groupId uint64) ([]rbac.Menu, error) {
	// 校验操作者权限
	ok, err := env.MenuAdapter.VerifyGroupIdWithOperator(operator, groupId)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("权限不足")
	}

	// 查询菜单
	return env.MenuAdapter.SelectMenuByGroupId(groupId)
}