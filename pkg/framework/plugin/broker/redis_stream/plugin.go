package broker_redis_stream

import (
	"context"
	"fmt"
	"framework/class/broker"
	"framework/class/logger"
	"framework/env"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

/*
	Redis Stream Broker
*/
var _ broker.Broker = &plugin{}

type plugin struct {
	opts        *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) Connect(dns string) error {
	client, err := connect(dns)
	if err != nil {
		return err
	}
	p.opts.redisClient = client
	return nil
}

func (p *plugin) Publish(topic string, message []byte) error {
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

		_, err := p.opts.redisClient.XAdd(xAddArgs).Result()
		return err
	}
}

func (p *plugin) Subscribe(topic string, handler broker.Handler) (broker.Subscriber, error) {
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
				res, err := p.opts.redisClient.XRead(xReadArgs).Result()
				if err != nil {
					if err == redis.Nil {
						time.Sleep(time.Millisecond * 100)
					} else {
						env.Logger.Log(logger.ErrorLevel, err)
					}
					continue
				}
				for _, stream := range res {
					for _, message := range stream.Messages {
						evt := &redisStreamEvent{
							header:      map[string]string{"Id": message.ID},
							topic:       stream.Stream,
							message:     nil,
							redisClient: p.opts.redisClient,
							err:         nil,
						}

						payload, ok := message.Values["Payload"].(string)
						if !ok {
							evt.err = fmt.Errorf("Assert Type field is failed\n")
							if err := handler(evt); err != nil {
								env.Logger.Log(logger.ErrorLevel, err)
							}
							continue
						}
						evt.message = []byte(payload)

						if err := handler(evt); err != nil {
							env.Logger.Log(logger.ErrorLevel, err)
							continue
						}
					}
				}
			}
		}
	}()

	return subscriber, nil
}
