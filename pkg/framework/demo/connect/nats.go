package main

import (
	"fmt"
	"framework/utils"
)

func main(){
	// root:123456@tcp(127.0.0.1:13306)
	dns := "nats://minigame:M123G321@192.168.0.20:8111"
	cli,err := utils.NatsClient(dns)
	if err != nil {
		panic(err)
	}
	fmt.Println(cli)
}
