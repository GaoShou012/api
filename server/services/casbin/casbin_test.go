package casbin

import (
	"api/config"
	"api/global"
	"github.com/casbin/casbin/v2"
	"testing"
)

func TestCasbinMysqlInit(t *testing.T) {
	config.LocalLoad()
	if err := global.InitCasbinAdapter(config.GetConfig().Casbin.DNS); err != nil {
		panic(err)
	}
	rbacFilePath := "./rbac_model.conf"

	e, err := casbin.NewEnforcer(rbacFilePath, global.CasbinAdapter)
	if err != nil {
		panic(err)
	}
	if err := e.LoadPolicy(); err != nil {
		panic(err)
	}
}
