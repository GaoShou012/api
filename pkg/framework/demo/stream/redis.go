package main

import (
	"encoding/json"
	"fmt"
	stream_redis_stream "framework/plugin/stream/redis_stream"
	"time"
)

func main() {
	dns := fmt.Sprintf("redis://:@127.0.0.1:17001?Db=0&PoolMax=100&PoolMin=10")

	topic := "testing_stream1"

	s := stream_redis_stream.New()
	if err := s.Connect(dns); err != nil {
		panic(err)
	}

	{
		m := make(map[string]string)
		m["Content"] = "123"
		m["ContentType"] = "text"
		m["Time"] = time.Now().String()
		j, err := json.Marshal(m)
		if err != nil {
			panic(err)
		}
		msgId, err := s.Push(topic, j)
		if err != nil {
			panic(err)
		}
		fmt.Println("message id:", msgId)

		event, err := s.PullById(topic, msgId)
		if err != nil {
			panic(err)
		}
		if event == nil {
			panic("message id is not exists")
		}
		fmt.Println("message:", event)
	}

	{
		events, err := s.Pull(topic, "0", 10)
		if err != nil {
			panic(err)
		}
		for _, event := range events {
			m := make(map[string]string)
			if err := json.Unmarshal(event.Message(), &m); err != nil {
				panic(err)
			}
			fmt.Println(m)

			// the ack to delete the event in stream.
			//event.Ack()
		}
	}

	{

	}
}
