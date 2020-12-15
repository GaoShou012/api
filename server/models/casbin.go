package models

import (
	"api/global"
	"github.com/casbin/casbin/v2"
	"regexp"
	"strings"
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

func (m *CasbinRule) NewCasbin() (*casbin.Enforcer, error) {
	//config.LocalLoad()
	//if err := initialize.InitCasbinAdapter(config.GetConfig().Casbin.DNS); err != nil {
	//	return nil, err
	//}
	//e, err := casbin.NewEnforcer(config.GetConfig().Casbin.RBACModelPath, global.CasbinAdapter)
	//if err != nil {
	//	return nil, err
	//}
	//if err := e.LoadPolicy(); err != nil {
	//	return nil, err
	//}
	//return e, nil
	return nil, nil
}

/**
添加casbin_rule表的规则其中
rules格式为
[][]string(AuthorityId,AuthorityId,AuthorityId)
*/
func (m *CasbinRule) AddCasbinPolicy(rules [][]string) error {
	e, err := m.NewCasbin()
	if err != nil {
		return err
	}
	success, err := e.AddPolicies(rules)
	if success != true {
		return err
	}
	return nil
}

/**
删除权限
*/
func (m *CasbinRule) RemoveCasbinPolicy(v int, p ...string) error {
	e, err := m.NewCasbin()
	if err != nil {
		return err
	}
	su, err := e.RemoveFilteredPolicy(v, p...)
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
	e, err := m.NewCasbin()
	if err != nil {
		return nil, err
	}
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps, nil
}

//判断权限
func (m *CasbinRule) ExecutePermission(sub interface{},obj interface{},act interface{}) (bool,error) {
	e, err := m.NewCasbin()
	if err != nil {
		return false,err
	}
<<<<<<< HEAD
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
=======
	sub := "bxcb" //v1
	obj := "POST" //v2
	act := "u"    //v0
>>>>>>> a4d4ef705b58083b43d0abd91c762e185f1e5ef7
	success, _ := e.Enforce(sub, obj, act)
	return success,err
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


//@function: ParamsMatch
//@description: 自定义规则函数
//@param: fullNameKey1 string, key2 string
//@return: bool

// RegexMatch determines whether key1 matches the pattern of key2 in regular expression.
func RegexMatch(key1 string, key2 string) bool {
	res, err := regexp.MatchString(key2, key1)
	if err != nil {
		panic(err)
	}
	return res
}

// KeyMatch2 determines whether key1 matches the pattern of key2 (similar to RESTful path), key2 can contain a *.
// For example, "/foo/bar" matches "/foo/*", "/resource1" matches "/:resource"
func KeyMatch2(key1 string, key2 string) bool {
	key2 = strings.Replace(key2, "/*", "/.*", -1)

	re := regexp.MustCompile(`(.*):[^/]+(.*)`)
	for {
		if !strings.Contains(key2, "/:") {
			break
		}

		key2 = re.ReplaceAllString(key2, "$1[^/]+$2")
	}

	return RegexMatch(key1, "^"+key2+"$")
}

func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return KeyMatch2(key1, key2)
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ParamsMatchFunc
//@description: 自定义规则函数
//@param: args ...interface{}
//@return: interface{}, error

func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return ParamsMatch(name1, name2), nil
}