package rbac_mysql_redis

import (
	"errors"
	"framework/class/rbac"
	"framework/env"
	"strconv"
	"strings"
)

var _ rbac.RBAC = &plugin{}

type plugin struct {
	opts *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) CreateApi(operator rbac.Operator, api rbac.Api) error {
	return env.ApiAdapter.Create(api)
}

func (p *plugin) DeleteApi(operator rbac.Operator, apiId uint64) (bool,error) {
	// 校验操作者权限
	{
		ok, err := env.ApiAdapter.Authority(operator, apiId)
		if err != nil {
			return false,err
		}
		if !ok {
			return false,errors.New("目标API权限不足")
		}
	}

	return env.ApiAdapter.Delete(apiId)
}

func (p *plugin) UpdateApi(operator rbac.Operator, apiId uint64, api rbac.Api) error {
	// 校验操作者权限
	{
		ok, err := env.ApiAdapter.Authority(operator, apiId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("目标API权限不足")
		}
	}

	return env.ApiAdapter.Update(apiId, api)
}

func (p *plugin) CreateMenu(operator rbac.Operator, groupId uint64, menu rbac.Menu) error {
	{
		ok, err := env.MenuAdapter.AuthorityMenuGroup(operator, groupId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("权限不足")
		}
	}

	return env.MenuAdapter.CreateMenu(menu)
}

func (p *plugin) DeleteMenu(operator rbac.Operator, menuId uint64) (bool,error) {
	{
		ok, err := env.MenuAdapter.AuthorityMenu(operator, menuId)
		if err != nil {
			return false,err
		}
		if !ok {
			return false,errors.New("权限不足")
		}
	}
	return env.MenuAdapter.DeleteMenu(menuId)
}

func (p *plugin) UpdateMenu(operator rbac.Operator, menuId uint64, menu rbac.Menu) error {
	// 校验操作者权限
	{
		ok, err := env.MenuAdapter.AuthorityMenu(operator, menuId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("")
		}
	}

	return env.MenuAdapter.UpdateMenu(menuId, menu)
}

func (p *plugin) CreateMenuGroup(operator rbac.Operator, group rbac.MenuGroup) error {
	return env.MenuAdapter.CreateMenuGroup(group)
}

func (p *plugin) DeleteMenuGroup(operator rbac.Operator, menuGroupId uint64) (bool, error) {
	{
		ok, err := env.MenuAdapter.AuthorityMenuGroup(operator, menuGroupId)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, errors.New("权限不足")
		}
	}

	return env.MenuAdapter.DeleteMenuGroup(menuGroupId)
}

func (p *plugin) UpdateMenuGroup(operator rbac.Operator, menuGroupId uint64, group rbac.MenuGroup) error {
	// 校验操作者权限
	{
		ok, err := env.MenuAdapter.AuthorityMenuGroup(operator, menuGroupId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("")
		}
	}

	return env.MenuAdapter.UpdateMenuGroup(menuGroupId, group)
}

func (p *plugin) CreateRole(operator rbac.Operator, role rbac.Role) error {
	return env.RoleAdapter.CreateRole(role)
}

func (p *plugin) DeleteRole(operator rbac.Operator, roleId uint64) (bool,error) {
	ok, err := env.RoleAdapter.Authority(operator, roleId)
	if err != nil {
		return false,err
	}
	if !ok {
		return false,errors.New("权限不足")
	}
	return env.RoleAdapter.DeleteRole(roleId)
}

func (p *plugin) UpdateRole(operator rbac.Operator, roleId uint64, role rbac.Role) error {
	ok, err := env.RoleAdapter.Authority(operator, roleId)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("权限不足")
	}
	return env.RoleAdapter.UpdateRole(roleId, role)
}

func (p *plugin) RoleAssocApi(operator rbac.Operator, roleId uint64, apiId uint64) error {
	// 校验操作者权限
	{
		ok, err := env.RoleAdapter.Authority(operator, roleId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("目标角色ID，权限不足")
		}
	}
	{
		ok, err := env.ApiAdapter.Authority(operator, apiId)
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
	api, err := env.ApiAdapter.SelectById(apiId)
	if err != nil {
		return err
	}

	// 角色关联API
	if err := env.RoleAdapter.AssocApi(role, api); err != nil {
		return err
	}

	return nil
}

func (p *plugin) RoleAssocMenu(operator rbac.Operator, roleId uint64, menuId uint64) error {
	// 校验操作者权限
	{
		ok, err := env.RoleAdapter.Authority(operator, roleId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("权限不足")
		}
	}
	{
		ok, err := env.MenuAdapter.AuthorityMenu(operator, menuId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("权限不足")
		}
	}

	// 查询角色
	role, err := env.RoleAdapter.SelectById(roleId)
	if err != nil {
		return err
	}

	// 查询菜单
	menu, err := env.MenuAdapter.SelectMenuById(menuId)
	if err != nil {
		return err
	}

	// 关联
	if err := env.RoleAdapter.AssocMenu(role, menu); err != nil {
		return err
	}

	return nil
}

func (p *plugin) RoleAssocMenuGroup(operator rbac.Operator, roleId uint64, menuGroupId uint64) error {
	{
		ok, err := env.RoleAdapter.Authority(operator, roleId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("权限不足")
		}
	}
	{
		ok, err := env.MenuAdapter.AuthorityMenuGroup(operator, menuGroupId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("权限不足")
		}
	}

	// 查询角色
	role, err := env.RoleAdapter.SelectById(roleId)
	if err != nil {
		return err
	}

	// 查询菜单
	group, err := env.MenuAdapter.SelectMenuGroupById(menuGroupId)
	if err != nil {
		return err
	}

	// 关联操作
	return env.RoleAdapter.AssocMenuGroup(role, group)
}

/*
	鉴权
*/
func (p *plugin) Enforcer(authorityId string, method string, path string) (bool, error) {
	arr := strings.Split(authorityId, ",")
	for _, roleId := range arr {
		num, err := strconv.Atoi(roleId)
		if err != nil {
			return false, err
		}
		ok, err := env.RoleAdapter.EnforcerApi(uint64(num), method, path)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}
	return false, nil
}
