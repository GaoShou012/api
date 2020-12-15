package initialize

import (
	"api/global"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"net/url"
	"strings"
)

//func InitCasbinAdapter(dns string) error {
//	u, err := url.Parse(dns)
//	if err != nil {
//		return err
//	}
//	databaseType := u.Scheme
//	username := u.User.Username()
//	password, _ := u.User.Password()
//	_dns := fmt.Sprintf("%s:%s@(%s)%s?%s", username, password, u.Host, u.Path, u.RawQuery)
//	a, err := gormadapter.NewAdapter(databaseType, _dns, true)
//	if err != nil {
//		return err
//	}
//	global.CasbinAdapter = a
//	return nil
//}

func InitCasbinEnforcer(dns string, path string) error {
	u, err := url.Parse(dns)
	if err != nil {
		return err
	}
	databaseType := u.Scheme
	username := u.User.Username()
	password, _ := u.User.Password()
	_dns := fmt.Sprintf("%s:%s@(%s)%s?%s", username, password, u.Host, u.Path, u.RawQuery)
	a, err := gormadapter.NewAdapter(databaseType, _dns, true)
	if err != nil {
		return err
	}
	e, err := casbin.NewEnforcer(path, a)
	if err != nil {
		return err
	}
	if err := e.LoadPolicy(); err != nil {
		return err
	}

	e.AddFunction("ParamsMatch", paramsMatchFunc)

	global.CasbinEnforcer = e

	return nil
}

//@author: GaoShou
//@function: ParamsMatchFunc
//@description: 自定义规则函数
//@param: args ...interface{}
//@return: interface{}, error

func paramsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return paramsMatch(name1, name2), nil
}

//@author: GaoShou
//@function: ParamsMatch
//@description: 自定义规则函数
//@param: fullNameKey1 string, key2 string
//@return: bool

func paramsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}
