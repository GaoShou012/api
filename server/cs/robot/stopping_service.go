package robot

import (
	"cs/env"
	"fmt"
	"framework/class/countdown"
	countdown_context "framework/plugin/countdown"
)

var StoppingServiceAgent *StoppingService

type StoppingService struct {
	merchantCode string
	sessionId    string
	countdownS1  countdown.Countdown
	countdownS2  countdown.Countdown
	*Callback
}

func (p *StoppingService) Init(merchantCode string, sessionId string, callback *Callback) {
	p.merchantCode = merchantCode
	p.sessionId = sessionId
	p.Callback = callback
	p.countdownS1 = countdown_context.New()
	p.countdownS2 = countdown_context.New()
}

func (p *StoppingService) Start() error {
	timeout, err := p.Callback.GetTimeoutOfStoppingServiceS1(p.merchantCode)
	if err != nil {
		return err
	}

	p.countdownS1.SetTimeoutCallback(timeout, func(counter uint64, args ...interface{}) {
		timeout, err := p.Callback.GetTimeoutOfStoppingServiceS2(p.merchantCode)
		if err != nil {
			env.Logger.Error(err)
			return
		}
		p.countdownS2.SetTimeoutCallback(timeout, func(counter uint64, args ...interface{}) {
			// 直接关闭会话
			p.Callback.OnCloseSession(p.merchantCode, p.sessionId)
		})
		p.countdownS2.Enable()
	})
	p.countdownS1.Enable()

	return nil
}

func (p *StoppingService) OnEvent(evt Event) {
	switch evt.Type() {
	case EventTypeVisitorMessage:
		p.Stop()
		SetSessionState(evt.SessionId(), SessionStateRobotService)
		ForwardEvent(evt)
		break
	default:
		err := fmt.Sprintf("没有相应的事件处理:%d", EventTypeVisitorMessage)
		env.Logger.Error(err)
		break
	}
}

func (p *StoppingService) Stop() error {
	p.countdownS1.Disable()
	p.countdownS2.Disable()
	return nil
}
