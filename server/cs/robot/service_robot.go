package robot

import (
	"cs/env"
	"fmt"
	"framework/class/countdown"
)

var _ Service = &ServiceRobot{}

type ServiceRobot struct {
	callback    *CallbackOfRobotService
	step        map[string]int8
	countdownS1 map[string]countdown.Countdown
	countdownS2 map[string]countdown.Countdown
}

func (p *ServiceRobot) OnInit(callback *Callback) error {
	p.step = make(map[string]int8)
	p.callback = callback.CallbackOfRobotService
	p.countdownS1 = make(map[string]countdown.Countdown)
	p.countdownS2 = make(map[string]countdown.Countdown)
	return nil
}

func (p *ServiceRobot) OnEntry(evt Event) {
	SetSessionState(evt.GetSessionId(), SessionStateRobotService)
	p.onEvent(evt)
}

func (p *ServiceRobot) onCountdownEvent(event countdown.Event) {
	evt := event.GetParams()[0].(Event)
	p.onEvent(evt)
}

func (p *ServiceRobot) onEvent(evt Event) {
Loop:
	switch p.step[evt.GetSessionId()] {
	case 0:
		// 关闭所有倒计时
		{
			ct, ok := p.countdownS1[evt.GetSessionId()]
			if ok {
				ct.Disable()
			}
		}
		{
			ct, ok := p.countdownS2[evt.GetSessionId()]
			if ok {
				ct.Disable()
			}
		}

		// 创建s1倒计时
		{
			ct, ok := p.countdownS1[evt.GetSessionId()]
			if !ok {
				ct = NewCountdown()
				p.countdownS1[evt.GetSessionId()] = ct
			}
		}

		// 配置s1倒计时
		{
			timeout, err := p.callback.S1Timeout(evt.GetMerchantCode())
			if err != nil {
				env.Logger.Error(err)
				return
			}
			ct := p.countdownS1[evt.GetSessionId()]
			ct.Config(timeout, p.onCountdownEvent, evt)
			ct.Enable()
		}

		p.step[evt.GetSessionId()]++
		break
	case 1:
		p.callback.S1OnTimeoutCall(evt)
		p.step[evt.GetSessionId()]++
		goto Loop
	case 2:
		// 关闭所有倒计时
		{
			ct, ok := p.countdownS1[evt.GetSessionId()]
			if ok {
				ct.Disable()
			}
		}
		{
			ct, ok := p.countdownS2[evt.GetSessionId()]
			if ok {
				ct.Disable()
			}
		}

		// 创建s2倒计时
		{
			ct, ok := p.countdownS2[evt.GetSessionId()]
			if !ok {
				ct = NewCountdown()
				p.countdownS2[evt.GetSessionId()] = ct
			}
		}

		// 配置s2倒计时
		{
			timeout, err := p.callback.S2Timeout(evt.GetMerchantCode())
			if err != nil {
				env.Logger.Error(err)
				return
			}
			ct := p.countdownS2[evt.GetSessionId()]
			ct.Config(timeout, p.onCountdownEvent, evt)
			ct.Enable()
		}

		p.step[evt.GetSessionId()]++
		break
	case 3:
		p.OnExit(evt)
		ServiceStoppingAgent.OnEntry(evt)
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

func (p *ServiceRobot) OnExit(evt Event) {
	{
		ct, ok := p.countdownS1[evt.GetSessionId()]
		if ok {
			ct.Disable()
		}
	}
	{
		ct, ok := p.countdownS2[evt.GetSessionId()]
		if ok {
			ct.Disable()
		}
	}
	delete(p.countdownS1, evt.GetSessionId())
	delete(p.countdownS2, evt.GetSessionId())
}

func (p *ServiceRobot) OnClean(sessionId string) {
	delete(p.step, sessionId)
	delete(p.countdownS1, sessionId)
	delete(p.countdownS2, sessionId)
}
