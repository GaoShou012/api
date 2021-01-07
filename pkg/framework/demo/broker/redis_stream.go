package main

import (
	"encoding/json"
	"fmt"
	"framework/class/broker"
	broker_redis_stream "framework/plugin/broker/redis_stream"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	dns := fmt.Sprintf("redis://:@127.0.0.1:17001?Db=0&PoolMax=100&PoolMin=10")
	b := broker_redis_stream.New()
	if err := b.Connect(dns); err != nil {
		panic(err)
	}

	// 监听数据
	_, err := b.Subscribe("testing_stream", func(evt broker.Event) error {
		//fmt.Println("message :", evt.Message())
		fmt.Println("header:", evt.Header())
		m := make(map[string]string)
		if err := json.Unmarshal(evt.Message(), &m); err != nil {
			panic(err)
		}
		fmt.Println("received:", m)
		return evt.Ack()
	})
	if err != nil {
		panic(err)
	}

	// 定时发送数据
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			<-ticker.C
			msg := make(map[string]string)
			msg["Type"] = "123"
			msg["Body"] = "3333"
			msg["Time"] = time.Now().String()
			j, err := json.Marshal(msg)
			if err != nil {
				panic(err)
			}
			b.Publish("testing_stream", j)
		}
	}()

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
