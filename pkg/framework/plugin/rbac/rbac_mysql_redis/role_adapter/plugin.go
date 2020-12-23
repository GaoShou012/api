package role_adapter

import (
	lib_model "framework/class/libs/model"
	"framework/class/rbac"
	"github.com/jinzhu/gorm"
)

var _ rbac.RoleAdapter = &plugin{}

type plugin struct {
	roleModel               rbac.Model
	roleAssocApiModel       rbac.Model
	roleAssocMenuGroupModel rbac.Model
	roleAssocMenuModel      rbac.Model
	*Callback
	dbMaster *gorm.DB
	dbSlave  *gorm.DB
	opts     *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) Authority(operator rbac.Operator, roleId uint64) (bool, error) {
	return p.Callback.Authority(operator, roleId)
}

func (p *plugin) CreateRole(role rbac.Role) error {
	res := p.dbMaster.Table(p.roleModel.GetTableName()).Create(role)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *plugin) DeleteRole(roleId uint64) (bool, error) {
	res := p.dbMaster.Table(p.roleModel.GetTableName()).Where("id=?", roleId).Delete(p.roleModel)
	if res.Error != nil {
		return false,res.Error
	}
	return true,nil
}

func (p *plugin) UpdateRole(roleId uint64, role rbac.Role) error {
	res := p.dbMaster.Table(p.roleModel.GetTableName()).Where("id=?", roleId).Updates(p.roleModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *plugin) SelectById(roleId uint64) (rbac.Role, error) {
	newModel := lib_model.NewModel(p.roleModel)
	res := p.dbSlave.Table(p.roleModel.GetTableName()).Where("id=?", roleId).Find(newModel)
	if res.Error != nil {
		return nil, res.Error
	}
	return newModel.(rbac.Role), nil
}

func (p *plugin) AssocMenuGroup(role rbac.Role, group rbac.MenuGroup) error {
	panic("implement me")
}

func (p *plugin) DisassociateMenuGroup(roleId, menuGroupId uint64) (bool, error) {
	panic("implement me")
}

func (p *plugin) AssocMenu(role rbac.Role, menu rbac.Menu) error {
	panic("implement me")
}

func (p *plugin) DisassociateMenu(roleId uint64, menuId uint64) (bool, error) {
	panic("implement me")
}

func (p *plugin) AssocApi(role rbac.Role, api rbac.Api) error {
	panic("implement me")
}

func (p *plugin) DisassociateApi(roleId uint64, apiId uint64) (bool, error) {
	panic("implement me")
}

func (p *plugin) EnforcerApi(roleId uint64, method string, path string) (bool, error) {
	return p.Callback.ExistsByRoleIdReqMethodAndPath(roleId, method, path)
}
