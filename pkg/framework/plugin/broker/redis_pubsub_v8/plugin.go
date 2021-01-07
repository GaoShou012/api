package broker_redis_pubsub_v8

import (
	"context"
	"framework/class/broker"
	"framework/class/logger"
	"framework/env"
	"os"
	"sync"
)

/*
	Redis Pub Sub Broker
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
	// encode the message
	_, err := p.opts.redisClient.Publish(context.TODO(), topic, message).Result()
	return err
}
func (p *plugin) Subscribe(topic string, handler broker.Handler) (broker.Subscriber, error) {
	sub := p.opts.redisClient.Subscribe(context.TODO(), topic)

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
				env.Logger.Log(logger.ErrorLevel, msg, err)
				return
			}

			evt := &event{
				topic:   topic,
				message: []byte(msg.Payload),
				err:     nil,
			}

			// handle the event
			if err := handler(evt); err != nil {
				env.Logger.Log(logger.ErrorLevel, err)
				continue
			}
		}
	}()

	return subscriber, nil
}
