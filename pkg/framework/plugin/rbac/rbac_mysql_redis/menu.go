package rbac_mysql_redis

import (
	"framework/class/rbac"
	"github.com/jinzhu/gorm"
)

var _ rbac.MenuAdapter = &MenuAdapter{}

type MenuAdapter struct {
	db   *gorm.DB
	opts *Options
}

func (p *MenuAdapter) UpdateMenuGroup(groupId uint64, group rbac.MenuGroup) error {
	res := p.db.Model(group).Where("id=?",groupId).Updates(group)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *MenuAdapter) SelectMenuGroup(operator rbac.Operator) ([]rbac.MenuGroup, error) {
	//res :=p.db.Where("")
	panic("implement me")
}

func (p *MenuAdapter) SelectAllMenuGroup(tenantId uint64) ([]rbac.MenuGroup, error) {
	panic("implement me")
}

func (p *MenuAdapter) VerifyGroupIdWithOperator(operator rbac.Operator, groupId uint64) (bool, error) {
	panic("implement me")
}

func (p *MenuAdapter) SelectMenuGroupById(groupId uint64) (rbac.MenuGroup, error) {
	panic("implement me")
}

func (p *MenuAdapter) CreateMenu(groupId uint64, menu rbac.Menu) error {
	res := p.db.Model(menu).Create(menu)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *MenuAdapter) DeleteMenu(tenantId uint64, menuId uint64) error {
	panic("implement me")
}

func (p *MenuAdapter) UpdateMenu(menuId uint64, menu rbac.Menu) error {
	panic("implement me")
}

func (p *MenuAdapter) SelectByGroupId(tenantId uint64, groupId uint64) ([]rbac.Menu, error) {
	panic("implement me")
}

func (p *MenuAdapter) SelectMenuById(menuId uint64) (rbac.Menu, error) {
	panic("implement me")
}

func (p *MenuAdapter) SelectMenuByGroupId(groupId uint64) ([]rbac.Menu, error) {
	panic("implement me")
}

func (p *MenuAdapter) SelectByRoleId(tenantId uint64, roleId uint64) error {
	panic("implement me")
}

func (p *MenuAdapter)CreateMenuGroup(operator rbac.Operator, group rbac.MenuGroup) error{
	res := p.db.Model(group).Create(group)
	if res.Error != nil {
		return res.Error
	}
	return nil
}