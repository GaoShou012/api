package main

import (
	"api/config"
	"api/libs/connect"
	"api/services/admin_api"
	_ "github.com/go-sql-driver/mysql"
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
	conf.Load("./config/database.ini")
}

func main() {
	ginMode := config.GetConfig().Base.GinMode
	ginPort := config.GetConfig().Base.GinPort

	//连接数据库
	connect.InitDB()

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
