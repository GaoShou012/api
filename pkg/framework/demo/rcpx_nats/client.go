package main

import (
	"context"
	"fmt"

	//"github.com/smallnest/rpcx/client"
	//"github.com/rpcxio/rpcx-etcd/serverplugin"
	etcd_client "github.com/rpcxio/rpcx-etcd/client"
	"github.com/smallnest/rpcx/client"
	"log"
)

func main() {
	d, err := etcd_client.NewEtcdV3Discovery("/nats_rpcx", "test", []string{"192.168.0.20:2379"}, nil)
	if err != nil {
		log.Fatalln("etcd discovery is failed", err)
	}
	cli := client.NewXClient("test", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	fmt.Println(cli)
	var rsp interface{}
	req := "213123123"
	if err := cli.Call(context.TODO(), "min", req, rsp); err != nil {
		log.Fatalln("etcd call is failed", err)
	}
	fmt.Println(rsp)
}
