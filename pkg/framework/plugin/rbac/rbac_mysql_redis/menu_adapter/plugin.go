package api_adapter

import (
	lib_model "framework/class/libs/model"
	"framework/class/rbac"
	"github.com/jinzhu/gorm"
)

var _ rbac.MenuAdapter = &plugin{}

type plugin struct {
	menuModel      rbac.Model
	menuGroupModel rbac.Model
	*Callback
	dbMaster *gorm.DB
	dbSlave  *gorm.DB
	opts     *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) AuthorityMenu(operator rbac.Operator, menuId uint64) (bool, error) {
	return p.Callback.AuthorityMenuId(operator, menuId)
}

func (p *plugin) CreateMenuGroup(group rbac.MenuGroup) error {
	res := p.dbMaster.Table(p.menuGroupModel.GetTableName()).Create(group)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *plugin) DeleteMenuGroup(groupId uint64) (bool, error) {
	res := p.dbMaster.Table(p.menuGroupModel.GetTableName()).Where("id=?", groupId).Delete(p.menuGroupModel)
	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (p *plugin) UpdateMenuGroup(groupId uint64, group rbac.MenuGroup) error {
	res := p.dbMaster.Table(p.menuGroupModel.GetTableName()).Where("id=?", groupId).Updates(group)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *plugin) SelectMenuGroupById(menuGroupId uint64) (rbac.MenuGroup, error) {
	panic("implement me")
}

func (p *plugin) AuthorityMenuGroup(operator rbac.Operator, groupId uint64) (bool, error) {
	return p.Callback.AuthorityMenuGroupId(operator, groupId)
}

func (p *plugin) CreateMenu(menu rbac.Menu) error {
	res := p.dbMaster.Table(p.menuModel.GetTableName()).Create(menu)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *plugin) DeleteMenu(menuId uint64) error {
	res := p.dbMaster.Table(p.menuModel.GetTableName()).Where("id=?", menuId).Delete(p.menuModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *plugin) UpdateMenu(menuId uint64, menu rbac.Menu) error {
	res := p.dbMaster.Table(p.menuModel.GetTableName()).Where("id=?", menuId).Updates(menu)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *plugin) SelectMenuById(menuId uint64) (rbac.Menu, error) {
	newModel := lib_model.NewModel(p.menuModel)
	res := p.dbSlave.Table(p.menuModel.GetTableName()).Where("id=?", menuId).Find(newModel)
	if res.Error != nil {
		return nil, res.Error
	}
	return newModel.(rbac.Menu), nil
}

func (p *plugin) SelectMenuGroup(operator rbac.Operator) ([]rbac.MenuGroup, error) {
	panic("implement me")
}

func (p *plugin) SelectMenuByGroupId(groupId uint64) ([]rbac.Menu, error) {
	panic("implement me")
}

func (p *plugin) DeleteMenuGroupById(groupId uint64) error {
	panic("implement me")
}

func (p *plugin) SelectByRoleId(roleId uint64) error {
	panic("implement me")
}
