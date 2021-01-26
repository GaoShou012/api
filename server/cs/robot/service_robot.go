package robot

import (
	"cs/env"
	"fmt"
	"framework/class/countdown"
)

var _ Service = &ServiceRobot{}

type ServiceRobot struct {
	robot       *Robot
	callback    *CallbackOfRobotService
	step        map[string]int8
	countdownS1 map[string]countdown.Countdown
	countdownS2 map[string]countdown.Countdown
}

func (p *ServiceRobot) OnInit(robot *Robot, callback *Callback) error {
	p.step = make(map[string]int8)
	p.callback = callback.CallbackOfRobotService
	p.countdownS1 = make(map[string]countdown.Countdown)
	p.countdownS2 = make(map[string]countdown.Countdown)
	return nil
}

func (p *ServiceRobot) OnEntry(evt Event) {
	p.robot.SetSessionStage(evt.GetSessionId(), SessionStageRobotServicing)
	p.callback.OnEntry(evt)
	p.step[evt.GetSessionId()] = 0
	p.onEvent(evt)
}

func (p *ServiceRobot) OnExit(evt Event) {
	{
		ct := p.countdownS1[evt.GetSessionId()]
		ct.Stop()
	}
	{
		ct := p.countdownS2[evt.GetSessionId()]
		ct.Stop()
	}
}

func (p *ServiceRobot) OnClean(sessionId string) {
	delete(p.step, sessionId)
	delete(p.countdownS1, sessionId)
	delete(p.countdownS2, sessionId)
}

func (p *ServiceRobot) onCountdownEvent(event countdown.Event) {
	evt := event.GetParams()[0].(Event)
	p.onEvent(evt)
}

func (p *ServiceRobot) onEvent(evt Event) {
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
		{
			_, ok := p.countdownS2[evt.GetSessionId()]
			if !ok {
				p.countdownS2[evt.GetSessionId()] = NewCountdown()
			}
		}
		goto Loop
	case 1:
		// 启动s1倒计时
		timeout, err := p.callback.S1Timeout(evt.GetMerchantCode())
		if err != nil {
			env.Logger.Error(err)
			return
		}
		ct := p.countdownS1[evt.GetSessionId()]
		ct.New(timeout, p.onCountdownEvent, evt)

		p.step[evt.GetSessionId()]++
		break
	case 2:
		p.callback.S1OnTimeoutCall(evt)
		p.step[evt.GetSessionId()]++
		goto Loop
	case 3:
		// 启动s2倒计时
		timeout, err := p.callback.S2Timeout(evt.GetMerchantCode())
		if err != nil {
			env.Logger.Error(err)
			return
		}
		ct := p.countdownS2[evt.GetSessionId()]
		ct.New(timeout, p.onCountdownEvent, evt)

		p.step[evt.GetSessionId()]++
		break
	case 4:
		p.OnExit(evt)
		AgentOfStoppingService.OnEntry(evt)
		break
	}
}

func (p *ServiceRobot) OnEvent(evt Event) {
	switch evt.GetType() {
	case EventTypeVisitorMessage:
		p.step[evt.GetSessionId()] = 0
		p.onEvent(evt)
		break
	default:
		err := fmt.Errorf("未知的事件类型,evt=%v", evt)
		env.Logger.Error(err)
	}
}
