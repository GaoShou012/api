package api_adapter

import (
	"framework/class/rbac"
	"github.com/jinzhu/gorm"
)

var _ rbac.RoleAdapter = &plugin{}

type plugin struct {
	roleModel               rbac.Role
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
	panic("implement me")
}

func (p *plugin) DeleteRole(roleId uint64) (bool, error) {
	panic("implement me")
}

func (p *plugin) UpdateRole(roleId uint64, role rbac.Role) error {
	panic("implement me")
}

func (p *plugin) SelectById(roleId uint64) (rbac.Role, error) {
	panic("implement me")
}

func (p *plugin) AssocMenuGroup(role rbac.Role, group rbac.MenuGroup) error {
	panic("implement me")
}

func (p *plugin) AssocMenu(role rbac.Role, menu rbac.Menu) error {
	panic("implement me")
}

func (p *plugin) AssocApi(role rbac.Role, api rbac.Api) error {
	panic("implement me")
}

func (p *plugin) EnforcerApi(roleId uint64, method string, path string) (bool, error) {
	panic("implement me")
}
