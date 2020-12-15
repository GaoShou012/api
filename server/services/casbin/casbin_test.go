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
<<<<<<< HEAD
	sub := "uid"
	obj := "/path/aa"
	act := "POST"
	success, _ := e.Enforce(sub, obj, act)
	fmt.Println(success)
=======
	fmt.Println("is ok:", ok)
>>>>>>> a4d4ef705b58083b43d0abd91c762e185f1e5ef7
}
