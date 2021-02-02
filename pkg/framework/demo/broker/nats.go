package main

import (
	"fmt"
	"framework/class/broker"
	broker_nats "framework/plugin/broker/nats"
	"framework/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	dns := "nats://minigame:M123G321@192.168.0.2:8111"
	natsClient, err := utils.NatsClient(dns)
	if err != nil {
		panic(err)
	}
	b := broker_nats.New(
		broker_nats.WithNatsClient(natsClient),
	)

	b.Subscribe("test", func(evt broker.Event) error {
		fmt.Println(evt.Message())
		return nil
	})

	if err := b.Publish("test",[]byte("123123")); err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		switch s := <-c; s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			//info := fmt.Errorf("got signal %s; stop server", s)
			//panic(info)
		case syscall.SIGHUP:
			//info := fmt.Errorf("got signal %s; go to deamon", s)
			continue
		}
		break
	}
}
