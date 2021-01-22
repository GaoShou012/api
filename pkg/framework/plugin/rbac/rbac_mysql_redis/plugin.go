package rbac_mysql_redis

import (
	"errors"
	"fmt"
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
	exists, err := p.opts.ApiAdapter.ExistsBydMethodAndPath(operator, api.GetMethod(), api.GetPath())
	if err != nil {
		return env.Logger.Error(err)
	}
	if exists {
		return fmt.Errorf("API请求已经存在，不能重复创建")
	}
	return p.opts.ApiAdapter.Create(api)
}

func (p *plugin) DeleteApi(operator rbac.Operator, apiId uint64) (bool, error) {
	// 校验操作者权限
	{
		ok, err := p.opts.ApiAdapter.Authority(operator, apiId)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, errors.New("目标API权限不足")
		}
	}

	return p.opts.ApiAdapter.Delete(apiId)
}

func (p *plugin) UpdateApi(operator rbac.Operator, apiId uint64, api rbac.Api) error {
	// 校验操作者权限
	{
		ok, err := p.opts.ApiAdapter.Authority(operator, apiId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("目标API权限不足")
		}
	}

	return p.opts.ApiAdapter.Update(apiId, api)
}

func (p *plugin) CreateMenu(operator rbac.Operator, groupId uint64, menu rbac.Menu) error {
	{
		ok, err := p.opts.MenuAdapter.AuthorityMenuGroup(operator, groupId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("权限不足")
		}
	}

	return p.opts.MenuAdapter.CreateMenu(menu)
}

func (p *plugin) DeleteMenu(operator rbac.Operator, menuId uint64) (bool, error) {
	{
		ok, err := p.opts.MenuAdapter.AuthorityMenu(operator, menuId)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, errors.New("权限不足")
		}
	}
	return p.opts.MenuAdapter.DeleteMenu(menuId)
}

func (p *plugin) UpdateMenu(operator rbac.Operator, menuId uint64, menu rbac.Menu) error {
	// 校验操作者权限
	{
		ok, err := p.opts.MenuAdapter.AuthorityMenu(operator, menuId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("")
		}
	}

	return p.opts.MenuAdapter.UpdateMenu(menuId, menu)
}

func (p *plugin) SelectMenuWithFieldsByMenuGroupId(operator rbac.Operator, menuGroupId uint64, fields string, out interface{}) error {
	panic("implement me")
}

func (p *plugin) SelectMenuWithFieldsByRoleIdMulti(operator rbac.Operator, roleIdMulti []uint64, fields string, out interface{}) error {
	return p.opts.MenuAdapter.SelectMenuWithFieldsByRoleIdMulti(operator, roleIdMulti, fields, out)
}

func (p *plugin) CreateMenuGroup(operator rbac.Operator, group rbac.MenuGroup) error {
	return p.opts.MenuAdapter.CreateMenuGroup(group)
}

func (p *plugin) DeleteMenuGroup(operator rbac.Operator, menuGroupId uint64) (bool, error) {
	{
		ok, err := p.opts.MenuAdapter.AuthorityMenuGroup(operator, menuGroupId)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, errors.New("权限不足")
		}
	}

	return p.opts.MenuAdapter.DeleteMenuGroup(menuGroupId)
}

func (p *plugin) UpdateMenuGroup(operator rbac.Operator, menuGroupId uint64, group rbac.MenuGroup) error {
	// 校验操作者权限
	{
		ok, err := p.opts.MenuAdapter.AuthorityMenuGroup(operator, menuGroupId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("")
		}
	}

	return p.opts.MenuAdapter.UpdateMenuGroup(menuGroupId, group)
}

func (p *plugin) SelectMenuGroupWithFieldsByRoleIdMulti(operator rbac.Operator, roleIdMulti []uint64, fields string, out interface{}) error {
	return p.opts.MenuAdapter.SelectMenuGroupWithFieldsByRoleIdMulti(operator, roleIdMulti, fields, out)
}

func (p *plugin) CreateRole(operator rbac.Operator, role rbac.Role) error {
	return p.opts.RoleAdapter.CreateRole(role)
}

func (p *plugin) DeleteRole(operator rbac.Operator, roleId uint64) (bool, error) {
	ok, err := p.opts.RoleAdapter.Authority(operator, roleId)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, errors.New("权限不足")
	}
	return p.opts.RoleAdapter.DeleteRole(roleId)
}

func (p *plugin) UpdateRole(operator rbac.Operator, roleId uint64, role rbac.Role) error {
	ok, err := p.opts.RoleAdapter.Authority(operator, roleId)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("权限不足")
	}
	return p.opts.RoleAdapter.UpdateRole(roleId, role)
}

func (p *plugin) RoleAssocApi(operator rbac.Operator, roleId uint64, apiId uint64) error {
	// 校验操作者权限
	{
		ok, err := p.opts.RoleAdapter.Authority(operator, roleId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("目标角色ID，权限不足")
		}
	}
	{
		ok, err := p.opts.ApiAdapter.Authority(operator, apiId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("目标API权限不足")
		}
	}

	// 查询角色
	role, err := p.opts.RoleAdapter.SelectById(roleId)
	if err != nil {
		return err
	}
	api, err := p.opts.ApiAdapter.SelectById(apiId)
	if err != nil {
		return err
	}

	// 角色关联API
	if err := p.opts.RoleAdapter.AssocApi(role, api); err != nil {
		return err
	}

	return nil
}

func (p *plugin) RoleDisassociateApi(operator rbac.Operator, assocId uint64) error {
	assoc, err := p.opts.RoleAdapter.SelectAssocApiById(assocId)
	if err != nil {
		return err
	}

	{
		ok, err := p.opts.RoleAdapter.Authority(operator, assoc.GetRoleId())
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("角色操作权限不足")
		}
	}
	{
		ok, err := p.opts.ApiAdapter.Authority(operator, assoc.GetApiId())
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("API操作权限不足")
		}
	}

	{
		_, err := p.opts.RoleAdapter.DisassociateApi(assocId)
		if err != nil {
			return err
		}
		return nil
	}
}

func (p *plugin) RoleAssocMenu(operator rbac.Operator, roleId uint64, menuId uint64) error {
	// 校验操作者权限
	{
		ok, err := p.opts.RoleAdapter.Authority(operator, roleId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("权限不足")
		}
	}
	{
		ok, err := p.opts.MenuAdapter.AuthorityMenu(operator, menuId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("权限不足")
		}
	}

	// 查询角色
	role, err := p.opts.RoleAdapter.SelectById(roleId)
	if err != nil {
		return err
	}

	// 查询菜单
	menu, err := p.opts.MenuAdapter.SelectMenuById(menuId)
	if err != nil {
		return err
	}

	// 关联
	if err := p.opts.RoleAdapter.AssocMenu(role, menu); err != nil {
		return err
	}

	return nil
}

func (p *plugin) RoleDisassociateMenu(operator rbac.Operator, assocId uint64) error {
	assoc, err := p.opts.RoleAdapter.SelectAssocMenuById(assocId)
	if err != nil {
		return err
	}

	{
		ok, err := p.opts.RoleAdapter.Authority(operator, assoc.GetRoleId())
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("角色操作权限不足")
		}
	}
	{
		ok, err := p.opts.MenuAdapter.AuthorityMenu(operator, assoc.GetMenuId())
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("菜单操作权限不足")
		}
	}

	{
		_, err := p.opts.RoleAdapter.DisassociateMenu(assocId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *plugin) RoleAssocMenuGroup(operator rbac.Operator, roleId uint64, menuGroupId uint64) error {
	{
		ok, err := p.opts.RoleAdapter.Authority(operator, roleId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("权限不足")
		}
	}
	{
		ok, err := p.opts.MenuAdapter.AuthorityMenuGroup(operator, menuGroupId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("权限不足")
		}
	}

	// 查询角色
	role, err := p.opts.RoleAdapter.SelectById(roleId)
	if err != nil {
		return err
	}

	// 查询菜单
	group, err := p.opts.MenuAdapter.SelectMenuGroupById(menuGroupId)
	if err != nil {
		return err
	}

	// 关联操作
	return p.opts.RoleAdapter.AssocMenuGroup(role, group)
}

func (p *plugin) RoleDisassociateMenuGroup(operator rbac.Operator, assocId uint64) error {
	assoc, err := p.opts.RoleAdapter.SelectAssocMenuGroupById(assocId)
	if err != nil {
		return err
	}

	{
		ok, err := p.opts.RoleAdapter.Authority(operator, assoc.GetRoleId())
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("角色操作权限不足")
		}
	}

	{
		ok, err := p.opts.MenuAdapter.AuthorityMenuGroup(operator, assoc.GetMenuGroupId())
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("菜单组操作权限不足")
		}
	}

	{
		_, err := p.opts.RoleAdapter.DisassociateMenuGroup(assocId)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
	鉴权
*/
func (p *plugin) Enforcer(operator rbac.Operator, roles string, method string, path string) error {
	api, err := p.opts.ApiAdapter.SelectByMethodAndPath(operator, method, path)
	if err != nil {
		return env.Logger.Error(err,"查询API失败")
	}
	if api == nil {
		return fmt.Errorf("API不存在")
	}
	if api.GetEnable() == false {
		return fmt.Errorf("API没有启用")
	}

	arr := strings.Split(roles, ",")
	for _, roleId := range arr {
		num, err := strconv.Atoi(roleId)
		if err != nil {
			return env.Logger.Error(err)
		}
		if p.opts.RoleAdapter.EnforcerApi(uint64(num), api.GetId()) == true {
			return nil
		}
	}
	return fmt.Errorf("权限不足")
}
