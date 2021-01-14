package robot

import (
	"cs/env"
	"fmt"
	"framework/class/countdown"
	"sync"
)

var _ Service = &ServiceStarting{}

type ServiceStarting struct {
	robot       *Robot
	mutex       sync.Mutex
	step        map[string]int8
	countdownS1 map[string]countdown.Countdown
	countdownS2 map[string]countdown.Countdown
	callback    *CallbackOfStartingService
}

func (p *ServiceStarting) OnInit(robot *Robot, callback *Callback) error {
	p.robot = robot
	p.callback = callback.CallbackOfStartingService
	p.countdownS1 = make(map[string]countdown.Countdown)
	p.countdownS2 = make(map[string]countdown.Countdown)
	return nil
}

func (p *ServiceStarting) OnEntry(evt Event) {
	p.robot.SetSessionStage(evt.GetSessionId(), SessionStageStarting)
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
	p.step[evt.GetSessionId()] = 0
	p.onEvent(evt)
}

func (p *ServiceStarting) OnExit(evt Event) {
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
}

func (p *ServiceStarting) OnClean(sessionId string) {
	delete(p.step, sessionId)
	delete(p.countdownS1, sessionId)
	delete(p.countdownS2, sessionId)
}

func (p *ServiceStarting) onCountdownEvent(event countdown.Event) {
	evt := event.GetParams()[0].(Event)
	p.onEvent(evt)
}

func (p *ServiceStarting) onEvent(evt Event) {
Loop:
	switch p.step[evt.GetSessionId()] {
	case 0:
		p.callback.OnEntry(evt)
		p.step[evt.GetSessionId()]++
		goto Loop
	case 1:
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
	case 2:
		p.callback.S1TimeoutCall(evt)
		p.step[evt.GetSessionId()]++
		goto Loop
	case 3:
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
			timeout, err := p.callback.S2Timeout(evt.GetSessionId())
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
	case 4:
		p.callback.S2TimeoutCall(evt)
		p.step[evt.GetSessionId()]++
		goto Loop
	case 5:
		// 进入自动结束流程
		p.OnExit(evt)
		AgentOfStoppingService.OnEntry(evt)
		break
	}
}

func (p *ServiceStarting) OnEvent(evt Event) {
	switch evt.GetType() {
	case EventTypeNewSession:
		p.OnEntry(evt)
		break
	case EventTypeVisitorMessage:
		p.OnExit(evt)
		AgentOfRobotServicing.OnEntry(evt)
		break
	default:
		err := fmt.Errorf("未知的事件类型,evt=%v", evt)
		env.Logger.Error(err)
	}
}
