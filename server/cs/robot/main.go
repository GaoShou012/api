package robot

import (
	"fmt"
	"framework/class/broker"
	lib_countdown "framework/libs/countdown"
)

type EventOfSessionStateChange struct {
	SessionId    string
	SessionState SessionState
}

type Robot struct {
	broker.Broker
	lib_countdown.Countdown
	Callback
	sessionState map[string]SessionState
	forwardEvent chan Event
}

func Init() {

}

func Run() {

}

func OnEvent(event Event) error {

}

func (r *Robot) PushEvent(evt Event) error {
	data, err := encodeEvent(evt)
	if err != nil {
		return err
	}

	if err := r.Broker.Publish("robot", data); err != nil {
		return err
	}

	return nil
}

func (r *Robot) OnEvent(evt Event) error {
	state, ok := GetSessionState(evt.SessionId())
	if !ok {
		state = SessionStateStartingService
		SetSessionState(evt.SessionId(), state)
	}

	switch state {
	case SessionStateStartingService:
		if err := StartingServiceAgent.OnEvent(evt); err != nil {
			return err
		}
		break
	case SessionStateRobotService:
		if err := RobotServiceAgent.OnEvent(evt); err != nil {
			return err
		}
		break
	case SessionStateHumanService:
		break
	case SessionStateRating:
		break
	case SessionStateStopping:
		break
	default:
		return fmt.Errorf("未知的会话阶段")
	}
	return nil
}

func (r *Robot) Handler() {
	r.Broker.Subscribe("robot", func(evt broker.Event) error {
		defer evt.Ack()
		robotEvt, err := decodeEvent(evt.Message())
		if err != nil {
			return err
		}
		r.OnEvent(robotEvt)
		return nil
	})
}
