package rbac

import (
	"errors"
	"framework/class/rbac"
	"framework/env"
)

/*
	创建菜单组
*/
func CreateMenuGroup(operator rbac.Operator, group rbac.MenuGroup) error {
	if err := env.MenuAdapter.CreateMenuGroup(operator, group); err != nil {
		return err
	}
	return nil
}

/*
	更新菜单组
*/
func UpdateMenuGroup(operator rbac.Operator, groupId uint64, group rbac.MenuGroup) error {
	// 校验操作者权限
	ok, err := env.MenuAdapter.VerifyGroupIdWithOperator(operator, groupId)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("权限不足")
	}

	// 更新菜单
	if err := env.MenuAdapter.UpdateMenuGroup(groupId, group); err != nil {
		return err
	}

	return nil
}

/*
	查询菜单组
*/
func SelectMenuGroupById(operator rbac.Operator, groupId uint64) (rbac.MenuGroup, error) {
	// 校验操作者权限
	ok, err := env.MenuAdapter.VerifyGroupIdWithOperator(operator, groupId)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("权限不足")
	}

	// 查询菜单组
	return env.MenuAdapter.SelectMenuGroupById(groupId)
}

/*
	查询所有菜单组
*/
func SelectAllMenuGroup(operator rbac.Operator) ([]rbac.MenuGroup, error) {
	return env.MenuAdapter.SelectAllMenuGroup(operator.GetTenantId())
}

