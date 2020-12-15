package casbin

import (
	"api/config"
	"api/global"
	"api/initialize"
	"api/models"
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
	e.AddFunction("ParamsMatch", models.ParamsMatchFunc)

	if err := e.LoadPolicy(); err != nil {
		panic(err)
	}
	sub := "uid"
	obj := "/path/aa"
	act := "POST"
	success, _ := e.Enforce(sub, obj, act)
	fmt.Println(success)
}
