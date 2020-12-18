package main

import (
	"encoding/json"
	"fmt"
	stream_redis_stream "framework/plugin/stream/redis_stream"
	"time"
)

func main() {
	dns := fmt.Sprintf("redis://:@127.0.0.1:17001?Db=0&PoolMax=100&PoolMin=10")

	s := stream_redis_stream.NewStream()
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
		msgId, err := s.Push("testing_stream1", j)
		if err != nil {
			panic(err)
		}
		fmt.Println("message id:", msgId)
	}

	{
		events, err := s.Pull("testing_stream1", "0", 10)
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
			event.Ack()
		}
	}
}