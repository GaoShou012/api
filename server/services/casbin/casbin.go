package casbin

import (
	"github.com/casbin/casbin/v2"
	"regexp"
	"strings"
)

//
//import "api/utils"
//
//func CreateAuthority(role){
//	utils.IMysql.Master.Where("authority_id =?",)
//}

//@author: [piexlmax](https://github.com/piexlmax)
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

func CasbinMysqlInit(dns string, rbacFilePath string) error {
	//u, err := url.Parse(dns)
	//if err != nil {
	//	return err
	//}
	//databaseType := u.Scheme
	//username := u.User.Username()
	//password, _ := u.User.Password()
	//_dns := fmt.Sprintf("%s:%s@(%s)%s?%s", username, password, u.Host, u.Path, u.RawQuery)
	//a, err := gormadapter.NewAdapter(databaseType, _dns, true)
	//if err != nil {
	//	return err
	//}
	//e, err := casbin.NewEnforcer(rbacFilePath, a)
	//if err != nil {
	//	return err
	//}
	//if err := e.LoadPolicy(); err != nil {
	//	return err
	//}

	//e.AddFunction("ParamsMatch", ParamsMatchFunc)
	return nil
}

func CasbinMysql(dns string, rbacFilePath string) (*casbin.Enforcer, error) {

	//username := ""
	//password := ""
	//path := ""
	//database := ""
	//dns := fmt.Sprintf("%s:%s@(%s)/%s", username, password, path, database)
	//
	//{
	//	adapter, err := gormadapter.NewAdapter("mysql", dns, true)
	//	if err != nil {
	//		return nil, err
	//	}
	//	enforcer, err := casbin.NewEnforcer()
	//}
	return nil, nil
}
