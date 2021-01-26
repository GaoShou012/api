package robot

import (
	"cs/env"
	"fmt"
	"framework/class/countdown"
)

var _ Service = &ServiceStopping{}

type ServiceStopping struct {
	robot       *Robot
	step        map[string]int8
	countdownS1 map[string]countdown.Countdown
	callback    *CallbackOfStoppingService
}

func (p *ServiceStopping) OnInit(robot *Robot, callback *Callback) error {
	p.callback = callback.CallbackOfStoppingService
	p.countdownS1 = make(map[string]countdown.Countdown)
	return nil
}

func (p *ServiceStopping) OnEntry(evt Event) {
	p.robot.SetSessionStage(evt.GetSessionId(), SessionStageStarting)
	p.callback.OnEntry(evt)
	p.step[evt.GetSessionId()] = 0
	p.onEvent(evt)
}

func (p *ServiceStopping) OnExit(evt Event) {
	{
		ct := p.countdownS1[evt.GetSessionId()]
		ct.Stop()
	}
}

func (p *ServiceStopping) OnClean(sessionId string) {
	delete(p.step, sessionId)
	delete(p.countdownS1, sessionId)
}

func (p *ServiceStopping) onCountdownEvent(event countdown.Event) {
	evt := event.GetParams()[0].(Event)
	p.onEvent(evt)
}

func (p *ServiceStopping) onEvent(evt Event) {
Loop:
	switch p.step[evt.GetSessionId()] {
	case 0:
		p.callback.OnEntry(evt)
		p.step[evt.GetSessionId()]++
		{
			_, ok := p.countdownS1[evt.GetSessionId()]
			if !ok {
				p.countdownS1[evt.GetSessionId()] = NewCountdown()
			}
		}
		goto Loop
	case 1:
		p.callback.S1TimeoutCall(evt)
		p.OnExit(evt)
		break
	}
}

func (p *ServiceStopping) OnEvent(evt Event) {
	switch evt.GetType() {
	case EventTypeVisitorMessage:
		p.OnExit(evt)
		AgentOfRobotServicing.onEvent(evt)
		break
	default:
		err := fmt.Errorf("未知的事件类型,evt=%v", evt)
		env.Logger.Error(err)
		break
	}
}
