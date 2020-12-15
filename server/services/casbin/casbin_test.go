package casbin

import (
	"api/config"
	"api/global"
	"api/initialize"
	"fmt"
	"testing"
)

func TestCasbinMysqlInit(t *testing.T) {
	config.LocalLoad()
	if err := initialize.InitCasbinEnforcer(config.GetConfig().Casbin.DNS, config.GetConfig().RBACModelPath); err != nil {
		panic(err)
	}

	sub := "888"
	obj := "abc"
	act := "POST"

	ok, err := global.CasbinEnforcer.Enforce(sub, obj, act)
	if err != nil {
		panic(err)
	}
	fmt.Println("is ok:", ok)
}
