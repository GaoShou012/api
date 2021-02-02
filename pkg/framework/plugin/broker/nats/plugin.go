package broker_nats

import (
	"framework/class/broker"
	"framework/class/logger"
	"framework/env"
	"github.com/nats-io/nats.go"
)

/*
	Redis Pub Sub Broker
*/
var _ broker.Broker = &plugin{}

type plugin struct {
	opts *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) Connect(dns string) error {
	//client, err := connect(dns)
	//if err != nil {
	//	return err
	//}
	//p.opts.redisClient = client
	return nil
}

func (p *plugin) Publish(topic string, message []byte) error {
	// encode the message
	err := p.opts.natsClient.Publish(topic, message)
	return err
}
func (p *plugin) Subscribe(topic string, handler broker.Handler) (broker.Subscriber, error) {
	sub, err := p.opts.natsClient.Subscribe(topic, func(msg *nats.Msg) {
		evt := &event{
			topic:   topic,
			message: msg.Data,
			err:     nil,
		}
		// handle the event
		if err := handler(evt); err != nil {
			env.Logger.Log(logger.ErrorLevel, err)
		}
	})
	if err != nil {
		return nil, err
	}

	//wg := sync.WaitGroup{}
	subscriber := &subscriber{}
	subscriber.SetupUnSubscribe(func() {
		sub.Unsubscribe()
	})
	//
	//go func() {
	//	wg.Add(1)
	//	defer wg.Done()
	//	for {
	//
	//		msg, err := sub.ReceiveMessage()
	//		if err != nil {
	//			env.Logger.Log(logger.ErrorLevel, msg, err)
	//			return
	//		}
	//
	//		evt := &event{
	//			topic:   topic,
	//			message: []byte(msg.Payload),
	//			err:     nil,
	//		}
	//
	//		// handle the event
	//		if err := handler(evt); err != nil {
	//			env.Logger.Log(logger.ErrorLevel, err)
	//			continue
	//		}
	//	}
	//}()

	return subscriber, nil
}
