package robot_v1

import (
	"cs/class/robot"
	"cs/plugin/robot/robot_v1/alert_event"
	"cs/plugin/robot/robot_v1/queue_event"
	"framework/class/broker"
)

type Options struct {
	broker.Broker
	alertEventCallback *alert_event.Callback
	queueEventCallback *queue_event.Callback
}

type Option func(o *Options)

func New(opts ...Option) robot.Robot {
	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	p := &plugin{
		opts: options,
	}
	if err := p.Init(); err != nil {
		panic(err)
	}

	return p
}

func WithBroker(broker broker.Broker) Option {
	return func(o *Options) {
		o.Broker = broker
	}
}

func WithQueueEventCallback(callback *queue_event.Callback) Option {
	return func(o *Options) {
		o.queueEventCallback = callback
	}
}

func WithAlertEventCallback(callback *alert_event.Callback) Option {
	return func(o *Options) {
		o.alertEventCallback = callback
	}
}
