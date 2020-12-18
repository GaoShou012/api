package models

import (
	"api/global"
)

type CasbinRule struct {
	//Ptype       string `json:"ptype" gorm:"column:p_type"`
	//AuthorityId string `json:"rolename" gorm:"column:v0"`
	//Path        string `json:"path" gorm:"column:v1"`
	//Method      string `json:"method" gorm:"column:v2"`
	Id    *int64
	PType *string
	V0    *string
	V1    *string
	V2    *string
	V3    *string
	V4    *string
	V5    *string
}

type CasbinInfo struct {
	Path   string
	Method string
}

/**
添加casbin_rule表的规则其中
rules格式为
[][]string(AuthorityId,AuthorityId,AuthorityId)
*/
func (m *CasbinRule) AddCasbinPolicy(rules [][]string) error {

	success, err := global.CasbinEnforcer.AddPolicies(rules)
	if success != true {
		return err
	}
	return nil
}

/**
删除权限
*/
func (m *CasbinRule) RemoveCasbinPolicy(v int, p ...string) error {
	su, err := global.CasbinEnforcer.RemoveFilteredPolicy(v, p...)
	if su != true {
		return err
	}
	return nil
}

//@function: GetPolicyPathByAuthorityId
//@description: 获取权限列表
//@param: authorityId string
//@return: pathMaps []request.CasbinInfo

func (m *CasbinRule) GetPolicyPathByAuthorityId(authorityId string) (pathMaps []CasbinInfo, err error) {
	list := global.CasbinEnforcer.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps, nil
}


/**
当修改api的请求方式时调用此函数
例如将/api/api的GET请求方式修改为POST
调用此方法后全表的/api/api将改为POST
这样就达到了修改的效果
@UpdateCasbinApi
*/
func (m *CasbinRule) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.DBMaster.Table("casbin_rule").Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	})
	return err.Error
}
