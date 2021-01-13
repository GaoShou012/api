package robot

import (
	"cs/env"
	"fmt"
	"framework/class/countdown"
	countdown_context "framework/plugin/countdown"
)

var _ Service = &ServiceStopping{}

type ServiceStopping struct {
	countdownS1 map[string]countdown.Countdown
	*Callback
}

func (p *ServiceStopping) OnInit(callback *Callback) error {
	p.Callback = callback
	p.countdownS1 = make(map[string]countdown.Countdown)
	return nil
}

func (p *ServiceStopping) OnEntry(evt Event) {
	SetSessionState(evt.GetSessionId(), SessionStateStopping)

	p.Callback.StoppingServiceOnEntry(evt)
	ct := countdown_context.New()

	timeout, err := p.Callback.StoppingServiceGetTimeoutS1(evt)
	if err != nil {
		env.Logger.Error(err)
		return
	}
	p.countdownS1[evt.GetSessionId()] = ct
	ct.SetTimeoutCallback(timeout, func(counter uint64, args ...interface{}) {
		evt := args[0].(Event)
		p.Callback.StoppingServiceOnTimeoutS1(evt)
		p.OnExit(evt)
	}, evt)
	ct.Enable()
}

func (p *ServiceStopping) OnExit(evt Event) {
	{
		ct, ok := p.countdownS1[evt.GetSessionId()]
		if ok {
			ct.Disable()
		}
	}

	delete(p.countdownS1, evt.GetSessionId())
}

func (p *ServiceStopping) OnEvent(evt Event) {
	switch evt.GetType() {
	case EventTypeVisitorMessage:
		p.OnExit(evt)
		ServiceRobotAgent.OnEntry(evt)
		break
	default:
		err := fmt.Errorf("未知的事件类型,evt=%v", evt)
		env.Logger.Error(err)
		break
	}
}
