package models

import (
	"api/config"
	"api/initialize"
	"fmt"
	"testing"
)

func TestCasbin_AddCasbinPolicy(t *testing.T) {
	m := CasbinRule{}
	rules := [][]string{}
	rules = append(rules, []string{"cm.xxx", "cm.xxx", "cm.xxx"})

	_ = m.AddCasbinPolicy(rules)
	config.LocalLoad()
	if err := initialize.InitMysqlMaster(config.GetConfig().MysqlMaster); err != nil {
		panic(err)
	}
	_ = m.RemoveCasbinPolicy(0,"a")
	_ = m.UpdateCasbinApi("bcb", "bxcb", "POST", "POST")
	pa,_ := m.GetPolicyPathByAuthorityId("u")
	fmt.Println(pa)
	//success,_ := m.ExecutePermission()
	//fmt.Println(success)
}
