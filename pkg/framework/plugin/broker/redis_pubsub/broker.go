package broker_redis_pubsub

import (
	"context"
	"framework/class/broker"
	"framework/class/logger"
	"github.com/go-redis/redis/v8"
	"os"
	"sync"
)

/*
	Redis Pub Sub Broker
*/
var _ broker.Broker = &redisPubSub{}

type redisPubSub struct {
	redisClient *redis.Client
	opts        *Options
}

func (b *redisPubSub) Init() error {
	b.redisClient = b.opts.redisClient
	if b.opts.logger == nil {
		b.opts.logger = &iLogger{}
	}
	return nil
}

func (b *redisPubSub) Connect(dns string) error {
	client, err := connect(dns)
	if err != nil {
		return err
	}
	b.redisClient = client
	return nil
}

func (b *redisPubSub) Publish(topic string, message []byte) error {
	// encode the message
	_, err := b.redisClient.Publish(context.TODO(), topic, message).Result()
	return err
}
func (b *redisPubSub) Subscribe(topic string, handler broker.Handler) (broker.Subscriber, error) {
	sub := b.redisClient.Subscribe(context.TODO(), topic)

	wg := sync.WaitGroup{}
	subscriber := &subscriber{}
	subscriber.SetupUnSubscribe(func() {
		//if err := sub.Unsubscribe(topic); err != nil {
		//	glog.Errorln(err)
		//}
		//wg.Wait()
		os.Exit(1)
	})

	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			msg, err := sub.ReceiveMessage(context.TODO())
			if err != nil {
				b.opts.logger.Log(logger.ErrorLevel, msg, err)
				return
			}

			evt := &redisPubSubEvent{
				topic:   topic,
				message: []byte(msg.Payload),
				err:     nil,
			}

			// handle the event
			if err := handler(evt); err != nil {
				b.opts.logger.Log(logger.ErrorLevel, err)
				continue
			}
		}
	}()

	return subscriber, nil
}
