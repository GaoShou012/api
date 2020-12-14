package casbin

import (
	"api/config"
	"api/global"
	"api/initialize"
	"fmt"
	"github.com/casbin/casbin/v2"
	"testing"
)

func TestCasbinMysqlInit(t *testing.T) {
	config.LocalLoad()
	if err := initialize.InitCasbinAdapter(config.GetConfig().Casbin.DNS); err != nil {
		panic(err)
	}

	e, err := casbin.NewEnforcer(config.GetConfig().Casbin.RBACModelPath, global.CasbinAdapter)
	if err != nil {
		panic(err)
	}
	if err := e.LoadPolicy(); err != nil {
		panic(err)
	}

	e.AddFunction("ParamsMatch",)

	sub := 888
	obj := "abc"
	act := "POST"

	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		panic(err)
	}
	fmt.Println("is ok:", ok)
}
