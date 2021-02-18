package initialize

import (
	"api/config"
	"api/global"
	libs_log "api/libs/logs"
	"flag"
	"fmt"
	"os"
	"testing"
)

var (
	confPath = flag.String("config", "../config/resource/dev.ini", "config file path")
)

func TestGorm(t *testing.T) {
	flag.Parse()
	// 初始化 mysql master

	conf := config.GetConfig()
	conf.Load(*confPath)

	{
		conf := config.GetConfig().MysqlMaster
		if err := InitMysqlMaster(conf); err != nil {
			libs_log.Error(err)
			os.Exit(0)
		}
	}

	// 初始化 mysql slave
	{
		conf := config.GetConfig().MysqlSlave
		if err := InitMysqlSlave(conf); err != nil {
			libs_log.Error(err)
			os.Exit(0)
		}
	}
	res := global.DBMaster.Exec("select * from rbac_api ")
	fmt.Println(res.Attrs())
	fmt.Println(123)
}
