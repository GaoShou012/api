package role_adapter

import (
	"fmt"
	"framework/class/rbac"
	lib_model "framework/libs/model"
)

var _ rbac.RoleAdapter = &plugin{}

type plugin struct {
	opts *Options
}

func (p *plugin) SelectAssocApiById(assocId uint64) (rbac.RoleAssocApi, error) {
	newModel := lib_model.New(p.opts.roleAssocApiModel).(rbac.RoleAssocApi)
	res := p.opts.dbSlave.Table(p.opts.roleAssocApiModel.GetTableName()).Where("id=?", assocId).Find(newModel)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, fmt.Errorf("角色关联API（id:%d）不存在", assocId)
		} else {
			return nil, res.Error
		}
	}
	return newModel, nil
}
func (p *plugin) SelectAssocMenuById(assocId uint64) (rbac.RoleAssocMenu, error) {
	newModel := lib_model.New(p.opts.roleAssocMenuModel).(rbac.RoleAssocMenu)
	res := p.opts.dbSlave.Table(p.opts.roleAssocMenuModel.GetTableName()).Where("id=?", assocId).Find(newModel)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, fmt.Errorf("角色关联菜单(id:%d)不存在", assocId)
		} else {
			return nil, res.Error
		}
	}
	return newModel, nil
}

func (p *plugin) SelectAssocMenuGroupById(assocId uint64) (rbac.RoleAssocMenuGroup, error) {
	newModel := lib_model.New(p.opts.roleAssocMenuGroupModel).(rbac.RoleAssocMenuGroup)
	res := p.opts.dbSlave.Table(p.opts.roleAssocMenuGroupModel.GetTableName()).Where("id=?", assocId).Find(newModel)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, fmt.Errorf("角色关联菜单(id:%d)不存在", assocId)
		} else {
			return nil, res.Error
		}
	}
	return newModel, nil
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) Authority(operator rbac.Operator, roleId uint64) (bool, error) {
	return p.opts.Callback.Authority(operator, roleId)
}

func (p *plugin) CreateRole(role rbac.Role) error {
	res := p.opts.dbMaster.Table(p.opts.roleModel.GetTableName()).Create(role)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *plugin) DeleteRole(roleId uint64) (bool, error) {
	res := p.opts.dbMaster.Table(p.opts.roleModel.GetTableName()).Where("id=?", roleId).Delete(p.opts.roleModel)
	if res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

func (p *plugin) UpdateRole(roleId uint64, role rbac.Role) error {
	res := p.opts.dbMaster.Table(p.opts.roleModel.GetTableName()).Where("id=?", roleId).Updates(p.opts.roleModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *plugin) SelectById(roleId uint64) (rbac.Role, error) {
	newModel := lib_model.New(p.opts.roleModel).(rbac.Role)
	res := p.opts.dbSlave.Table(p.opts.roleModel.GetTableName()).Where("id=?", roleId).Find(newModel)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, fmt.Errorf("角色ID(%d)不存在", roleId)
		} else {
			return nil, res.Error
		}
	}
	return newModel, nil
}

func (p *plugin) AssocMenuGroup(role rbac.Role, group rbac.MenuGroup) error {
	model := p.opts.Callback.AssocMenuGroup(role, group)
	res := p.opts.dbMaster.Table(p.opts.roleAssocMenuModel.GetTableName()).Create(model)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *plugin) DisassociateMenuGroup(assocId uint64) (bool, error) {
	res := p.opts.dbMaster.Table(p.opts.roleAssocMenuGroupModel.GetTableName()).Where("id=?", assocId).Delete(nil)
	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (p *plugin) AssocMenu(role rbac.Role, menu rbac.Menu) error {
	model := p.opts.Callback.AssocMenu(role, menu)
	res := p.opts.dbMaster.Table(p.opts.roleAssocMenuModel.GetTableName()).Create(model)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *plugin) DisassociateMenu(assocId uint64) (bool, error) {
	res := p.opts.dbMaster.Table(p.opts.roleAssocMenuModel.GetTableName()).Where("id=?", assocId).Delete(nil)
	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (p *plugin) AssocApi(role rbac.Role, api rbac.Api) error {
	model := p.opts.Callback.AssocApi(role, api)
	res := p.opts.dbMaster.Table(p.opts.roleAssocApiModel.GetTableName()).Create(model)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *plugin) DisassociateApi(assocId uint64) (bool, error) {
	res := p.opts.dbMaster.Table(p.opts.roleAssocApiModel.GetTableName()).Where("id=?", assocId).Delete(nil)
	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (p *plugin) EnforcerApi(roleId uint64, method string, path string) (bool, error) {
	return p.opts.Callback.ExistsByRoleIdReqMethodAndPath(roleId, method, path)
}
