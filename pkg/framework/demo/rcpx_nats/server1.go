package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/rpcxio/rpcx-etcd/serverplugin"
	"github.com/smallnest/rpcx/server"
	"log"
	"time"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	fmt.Printf("i am call: %d * %d = %d\n", args.A, args.B, reply.C)
	return nil
}

func (t *Arith) Add(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	fmt.Printf("i am call: %d + %d = %d\n", args.A, args.B, reply.C)
	return nil
}

func (t *Arith) Say(ctx context.Context, args *string, reply *string) error {
	*reply = "hello " + *args
	return nil
}

var (
	addr = flag.String("addr", ":8972", "server address")
)

func main(){
	flag.Parse()
	s := server.NewServer()

	plugin := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@:8972",
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
	s.Plugins.Add(plugin)

	if err := s.RegisterName("Arith", new(Arith), ""); err != nil {
		panic(err)
	}
	if err :=  s.Serve("tcp", *addr); err != nil {
		panic(err)
	}
	//select {}
}
