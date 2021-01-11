package robot

import (
	"encoding/json"
	"fmt"
	"framework/class/countdown"
	countdown_context "framework/plugin/countdown"
)

type StartingService struct {
	newCountdown                  NewCountdownForVisitorDoesNotAskOfStartingService
	countdownForVisitorDoesNotAsk countdown.Countdown
	merchant                      Merchant
	Callback
}

func (p *StartingService) Init(merchant Merchant) {
	p.merchant = merchant
}

func (p *StartingService) OnEvent(evt Event) error {
	switch evt.Type() {
	case EventTypeNewSession:
		session := &EventOfNewSession{}
		if err := json.Unmarshal(evt.Data(), session); err != nil {
			return err
		}

		// 新服务开始时的回调
		p.Callback.NewStartingService(session)

		// 获取商户的配置
		timeout, err := p.Callback.GetTimeoutOfVisitorDoesNotAskOnStartingService(session.MerchantCode)
		if err != nil {
			return err
		}

		// 新建倒计时
		ct := countdown_context.New()
		ct.SetTimeoutCallback(timeout, func(counter uint64, args ...interface{}) {
			session := args[0].(*EventOfNewSession)
			p.Callback.TimeoutOfVisitorDoesNotAskOnStartingService(counter, session)
		}, session)

		// 启动倒计时
		ct.Enable()

		// 保存倒计时
		p.countdownForVisitorDoesNotAsk = ct
		break
	case EventTypeVisitorMessage:
		p.countdownForVisitorDoesNotAsk.Disable()
		SetSessionState(evt.SessionId(), SessionStateRobotService)
		ForwardEvent(evt)
		break
	default:
		err := fmt.Errorf("开场白阶段，没有事件处理:%d", EventTypeNewSession)
		return err
	}

	return nil
}
