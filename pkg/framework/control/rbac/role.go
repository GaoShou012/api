package rbac

import (
	"errors"
	"framework/class/rbac"
	"framework/env"
)

/*
	创建角色
*/
func CreateRole(operator rbac.Operator, role rbac.Role) error {
	if err := env.RoleAdapter.CreateRole(operator, role); err != nil {
		return err
	}

	return nil
}

/*
	更新角色
*/
func UpdateRole(operator rbac.Operator, roleId uint64, role rbac.Role) error {
	return nil
}

/*
	查询角色
*/
func SelectRoleById(operator rbac.Operator, roleId uint64) (rbac.Role, error) {
	return nil, nil
}

/*
	关联角色&菜单组
*/
func RoleAssocMenuGroup(operator rbac.Operator, roleId uint64, groupId uint64) error {
	// 校验操作者权限
	ok, err := env.MenuAdapter.VerifyGroupIdWithOperator(operator, groupId)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("权限不足")
	}

	// 查询角色
	role, err := env.RoleAdapter.SelectById(roleId)
	if err != nil {
		return err
	}

	// 查询菜单组
	group, err := env.MenuAdapter.SelectMenuGroupById(groupId)
	if err != nil {
		return err
	}

	// 关联
	if err := env.RoleAdapter.AssocMenuGroup(role, group); err != nil {
		return err
	}

	return nil
}

/*
	关联角色&API
*/
func RoleAssocApi(operator rbac.Operator, roleId uint64, apiId uint64) error {
	{
		ok, err := env.RoleAdapter.VerifyIdWithOperator(roleId, operator)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("目标角色ID，权限不足")
		}
	}

	{
		ok, err := env.ApiAdapter.VerifyIdWithOperator(apiId, operator)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("目标API权限不足")
		}
	}

	// 查询角色ID
	role, err := env.RoleAdapter.SelectById(roleId)
	if err != nil {
		return err
	}

	// 查询API
	api, err := env.ApiAdapter.SelectById(operator.GetTenantId(), apiId)
	if err != nil {
		return err
	}

	// 角色关联API
	if err := env.RoleAdapter.AssocApi(role, api); err != nil {
		return err
	}

	return nil
}
