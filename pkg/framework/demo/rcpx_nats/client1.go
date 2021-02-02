package main

import (
	"context"
	"flag"
	etcd_client "github.com/rpcxio/rpcx-etcd/client"
	"github.com/smallnest/rpcx/client"

	"log"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

var (
	addr       = flag.String("addr", "127.0.0.1:8972", "server address")
)

func main(){
	flag.Parse()

	d, err := etcd_client.NewEtcdV3Discovery("/nats_rpcx", "Arith", []string{"192.168.0.20:2379"}, nil)
	if err != nil {
		log.Fatalln("etcd discovery is failed", err)
	}

	//d,err := client.NewPeer2PeerDiscovery("tcp@" + *addr, "")
	//if err != nil {
	//	panic(err)
	//}
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &Args{
		A: 10,
		B: 20,
	}

	reply := &Reply{}
	err = xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
