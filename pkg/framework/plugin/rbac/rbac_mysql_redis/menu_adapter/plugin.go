package menu_adapter

import (
	"framework/class/rbac"
	lib_model "framework/libs/model"
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
	menuGroup := lib_model.NewModel(p.menuGroupModel)
	res := p.dbMaster.Table(p.menuGroupModel.GetTableName()).Where("id=?", menuGroupId).Find(menuGroup)
	if res.Error != nil {
		return nil, res.Error
	}
	return menuGroup.(rbac.MenuGroup), nil
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

func (p *plugin) DeleteMenu(menuId uint64) (bool, error) {
	res := p.dbMaster.Table(p.menuModel.GetTableName()).Where("id=?", menuId).Delete(p.menuModel)
	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {
		return false, nil
	} else {
		return true, nil
	}
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

//func (p *plugin) SelectMenuGroup(operator rbac.Operator) ([]rbac.MenuGroup, error) {
//
//	menuGroup := lib_model.NewModel(p.menuGroupModel)
//	res := p.dbMaster.Table(p.menuGroupModel.GetTableName()).Where("tenant_id=?", operator.GetTenantId()).Find(menuGroup)
//	if res.Error != nil {
//		return nil, res.Error
//	}
//	return menuGroup.([]rbac.MenuGroup), nil
//}

func (p *plugin) SelectMenuByGroupId(groupId uint64) ([]rbac.Menu, error) {

	menu := lib_model.NewModel(p.menuModel)
	res := p.dbMaster.Table(p.menuModel.GetTableName()).Where("group_id=?", groupId).Find(menu)
	if res.Error != nil {
		return nil, res.Error
	}
	return menu.([]rbac.Menu), nil
}

func (p *plugin) DeleteMenuGroupById(groupId uint64) error {
	menuGroup := lib_model.NewModel(p.menuGroupModel)
	res := p.dbMaster.Table(p.menuModel.GetTableName()).Where("group_id=?", groupId).Delete(menuGroup)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *plugin) SelectByRoleId(roleId uint64) error {

	return nil
}
