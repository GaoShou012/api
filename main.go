package main

import (
	"api/config"
	libs_log "api/libs/logs"
	"api/services/admin_api"
	"api/utils"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

var (
	confPath = flag.String("config", "./config/dev.ini", "config file path")
)

func init() {
	flag.Parse()
	conf := config.GetConfig()
	conf.Load(*confPath)

	{
		conf := config.GetConfig().MysqlMaster
		if err := utils.IMysql.InitMaster(&conf); err != nil {
			libs_log.Error(err)
			os.Exit(0)
		}
	}

	{
		conf := config.GetConfig().MysqlSlave
		if err := utils.IMysql.InitSlave(&conf); err != nil {
			libs_log.Error(err)
			os.Exit(0)
		}
	}

	{
		conf := config.GetConfig().Redis
		if err := utils.InitRedis(&conf); err != nil {
			libs_log.Error(err)
			os.Exit(0)
		}
	}
}

func main() {
	ginMode := config.GetConfig().Base.GinMode
	ginPort := config.GetConfig().Base.GinPort

	// 初始化Gin
	r := gin.New()
	gin.SetMode(ginMode)

	// admin api service
	{
		httpService := admin_api.HttpService{}
		httpService.Route(r)
	}

	// 启动服务
	go func() {
		if err := r.Run(fmt.Sprintf(":%d", ginPort)); err != nil {
			// 打印日志
			return
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		switch s := <-c; s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			//logs.Error("got signal %s; stop server", s)
		case syscall.SIGHUP:
			//logs.Error("got signal %s; go to deamon", s)
			continue
		}
		break
	}
}
