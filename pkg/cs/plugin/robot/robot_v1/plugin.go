package robot_v1

import (
	"cs/class/robot"
	"cs/env"
	"cs/meta"
	"cs/plugin/robot/robot_v1/alert_event"
	"cs/plugin/robot/robot_v1/queue_event"
	"errors"
	"fmt"
	"framework/class/broker"
	"framework/class/logger"
	"time"
)

var _ robot.Robot = &plugin{}

type plugin struct {
	debug              bool
	alertEventHandler  *alert_event.Handler
	queueEventCallback *queue_event.Callback
	broker             broker.Broker
	opts               *Options
}

func (p *plugin) Init() error {
	p.debug = true
	p.broker = p.opts.Broker

	// init queue_event
	if p.opts.queueEventCallback == nil {
		return errors.New("queue event callback is nil\n")
	}
	p.queueEventCallback = p.opts.queueEventCallback

	// init alert event
	p.alertEventHandler = &alert_event.Handler{}
	if err := p.alertEventHandler.Init(p.broker, p.opts.alertEventCallback); err != nil {
		return err
	}
	return nil
}

func (p *plugin) IsDebug() bool {
	if p.debug == false {
		return false
	}
	return env.Logger.IsDebug()
}

func (p *plugin) QueueNotification(session meta.Session, client meta.Client) error {
	if p.IsDebug() {
		info := fmt.Sprintf("sessionId:%s,尝试触发排队通知", session.GetSessionId())
		env.Logger.Log(logger.InfoLevel, info)
	}

	// flag one times
	{
		ok, err := env.Session.SetFlagNX(session.GetSessionId(), "QueueNotification", time.Now().String())
		if err != nil {
			return err
		}

		if p.IsDebug() {
			info := fmt.Sprintf("sessionId:%s,尝试触发排队通知,%v", session.GetSessionId(), ok)
			env.Logger.Log(logger.InfoLevel, info)
		}

		if !ok {
			return nil
		}
	}

	// publish warn message
	{
		message := p.queueEventCallback.Notification(session, client)
		if message == nil {
			return nil
		}
		if err := env.Gateway.Publish(client.GetUUID(), message); err != nil {
			return err
		}
	}

	return nil
}

func (p *plugin) AlertEvent(event robot.AlertEvent, timeout time.Duration, options map[string]string) error {
	return p.alertEventHandler.PublishEvent(event, timeout, options)
}

func (p *plugin) AlertEventHandler() error {
	p.alertEventHandler.Run()
	return nil
}
