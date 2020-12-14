package main

import (
	"api/config"
	"api/initialize"
	libs_ip_location "api/libs/ip_location"
	libs_log "api/libs/logs"
	"api/services/admin_api"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

var (
	confPath = flag.String("config", "./config/resource/dev.ini", "config file path")
)

func init() {
	flag.Parse()
	conf := config.GetConfig()
	conf.Load(*confPath)

	// 初始化 ip location
	{
		dbPath := config.GetConfig().IpLocation.Path
		if err := libs_ip_location.Init(dbPath); err != nil {
			libs_log.Error(err)
			os.Exit(0)
		}
	}

	// 初始化 casbin adapter
	{
		conf := config.GetConfig().Casbin
		if err := initialize.InitCasbinAdapter(conf.DNS); err != nil {
			libs_log.Error(err)
			os.Exit(0)
		}
	}

	// 初始化 mysql master
	{
		conf := config.GetConfig().MysqlMaster
		if err := initialize.InitMysqlMaster(conf); err != nil {
			libs_log.Error(err)
			os.Exit(0)
		}
	}

	// 初始化 mysql slave
	{
		conf := config.GetConfig().MysqlSlave
		if err := initialize.InitMysqlSlave(conf); err != nil {
			libs_log.Error(err)
			os.Exit(0)
		}
	}

	// 初始化 redis
	{
		conf := config.GetConfig().Redis
		if err := initialize.InitRedis(conf); err != nil {
			libs_log.Error(err)
			os.Exit(0)
		}
	}

	fmt.Println("初始化完成")
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
