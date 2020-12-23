package gateway_v1

import (
	"cs/class/gateway"
	"encoding/json"
	"errors"
	"framework/class/broker"
)

var _ gateway.Gateway = &plugin{}

type plugin struct {
	topic string
	opts  *Options
	broker.Broker
}

func (p *plugin) Init() error {
	if p.opts.topic != "" {
		p.topic = p.opts.topic
	} else {
		p.topic = "gateway"
	}

	if p.opts.Broker == nil {
		return errors.New("broker is nil")
	}
	p.Broker = p.opts.Broker

	return nil
}

func (p *plugin) Publish(uuid string, message []byte) error {
	evt := &event{
		T: gateway.MessageTypeText,
		H: map[string]string{"UUID": uuid},
		B: message,
	}

	j, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	if err := p.Broker.Publish(p.topic, j); err != nil {
		return err
	}

	return nil
}

func (p *plugin) PublishControl(control gateway.Control) error {
	evt := &event{
		T: gateway.MessageTypeText,
		H: map[string]string{"CtlType": control.CtlType()},
		B: nil,
	}

	{
		j, err := json.Marshal(control)
		if err != nil {
			return err
		}
		evt.B = j
	}

	{
		j, err := json.Marshal(evt)
		if err != nil {
			return err
		}
		if err := p.Broker.Publish(p.topic, j); err != nil {
			return err
		}
	}

	return nil
}

func (p *plugin) Subscribe(handler gateway.Handler) (broker.Subscriber, error) {
	return p.Broker.Subscribe(p.topic, func(evt broker.Event) error {
		msg := &event{}

		if err := json.Unmarshal(evt.Message(), msg); err != nil {
			return err
		}

		if err := handler(msg); err != nil {
			return err
		}

		return nil
	})
}
