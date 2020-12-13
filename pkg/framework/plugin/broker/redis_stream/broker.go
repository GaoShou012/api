package broker_redis_stream

import (
	"context"
	"fmt"
	"framework/class/broker"
	"framework/class/logger"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
	"time"
)

/*
	Redis Stream Broker
*/
var _ broker.Broker = &redisStream{}

type redisStream struct {
	redisClient *redis.Client
	opts        *Options
}

func (b *redisStream) Init() error {
	b.redisClient = b.opts.redisClient
	if b.opts.logger == nil {
		b.opts.logger = &iLogger{}
	}
	return nil
}

func (b *redisStream) Connect(dns string) error {
	client, err := connect(dns)
	if err != nil {
		return err
	}
	b.redisClient = client
	return nil
}

func (b *redisStream) Publish(topic string, message []byte) error {
	{
		values := make(map[string]interface{})
		values["Payload"] = string(message)
		xAddArgs := &redis.XAddArgs{
			Stream:       topic,
			MaxLen:       0,
			MaxLenApprox: 0,
			ID:           "*",
			Values:       values,
		}

		fmt.Println("values:", values)

		_, err := b.redisClient.XAdd(context.TODO(), xAddArgs).Result()
		return err
	}
}

func (b *redisStream) Subscribe(topic string, handler broker.Handler) (broker.Subscriber, error) {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	subscriber := &subscriber{}
	subscriber.SetupUnSubscribe(func() {
		cancel()
		wg.Wait()
	})

	go func() {
		wg.Add(1)
		defer wg.Done()

		xReadArgs := &redis.XReadArgs{
			Streams: []string{topic, "0"},
			Count:   1,
			Block:   -1,
		}
		for {
			select {
			case <-ctx.Done():
				return
			default:
				res, err := b.redisClient.XRead(context.TODO(), xReadArgs).Result()
				if err != nil {
					if err == redis.Nil {
						time.Sleep(time.Millisecond * 100)
					} else {
						b.opts.logger.Log(logger.ErrorLevel, err)
					}
					continue
				}
				for _, stream := range res {
					for _, message := range stream.Messages {
						evt := &redisStreamEvent{
							header:      map[string]string{"Id": message.ID},
							topic:       stream.Stream,
							message:     nil,
							redisClient: b.redisClient,
							err:         nil,
						}

						payload, ok := message.Values["Payload"].(string)
						if !ok {
							evt.err = fmt.Errorf("Assert Type field is failed\n")
							if err := handler(evt); err != nil {
								if b.opts.logger != nil {
									b.opts.logger.Log(logger.ErrorLevel, err)
								} else {
									log.Fatal(err)
								}
							}
							continue
						}
						evt.message = []byte(payload)

						if err := handler(evt); err != nil {
							b.opts.logger.Log(logger.ErrorLevel, err)
							continue
						}
					}
				}
			}
		}
	}()

	return subscriber, nil
}
