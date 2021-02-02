package main

import (
	"context"
	"fmt"
	"framework/utils"
	"github.com/rcrowley/go-metrics"
	"github.com/rpcxio/rpcx-etcd/serverplugin"
	"github.com/smallnest/rpcx/server"
	"log"
	"time"
)

type Service1 struct {
}

func (s *Service1) Min(ctx context.Context, req *string,rsp *string) error {
	fmt.Println("calling",req)
	rsp = req
	return nil
}

func main() {
	cli, err := utils.NatsClient("nats://minigame:M123G321@192.168.0.20:8111")
	if err != nil {
		panic(err)
	}
	fmt.Println(cli)

	svr := server.NewServer()
	plugin := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@192.168.0.111:9901",
		EtcdServers:    []string{"192.168.0.20:2379"},
		BasePath:       "/nats_rpcx",
		Metrics:        metrics.NewRegistry(),
		Services:       nil,
		UpdateInterval: time.Minute,
		Options:        nil,
	}
	if err := plugin.Start(); err != nil {
		log.Fatalln("plugin start:", err)
	}
	svr.Plugins.Add(plugin)

	if err := svr.RegisterName("test", new(Service1), ""); err != nil {
		log.Fatalln("register rpcx service failed", err)
	}
	//err = svr.RegisterFunctionName("test/login", "hello", func(ctx context.Context, req interface{}, rsp interface{}) error {
	//	fmt.Println(req)
	//	return nil
	//}, "")
	//if err != nil {
	//	log.Fatalln("注册rcpx服务失败", err)
	//}
	if err := svr.Serve("http", ":9901"); err != nil {
		log.Fatalln("启动server失败,", err)
	}
}
