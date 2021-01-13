package robot

import (
	"framework/class/countdown"
)

var _ Service = &StartingService{}

type StartingService struct {
	countdownForVisitorDoesNotAsk map[string]countdown.Countdown

	merchant Merchant
	Callback
}

func (p *StartingService) OnInit(robot *Robot, callback *Callback) error {
	panic("implement me")
}

func (p *StartingService) OnEntry(evt Event) {
	panic("implement me")
}

func (p *StartingService) OnEvent(evt Event) {
	panic("implement me")
}

func (p *StartingService) OnExit(evt Event) {
	panic("implement me")
}

func (p *StartingService) OnClean(sessionId string) {
	panic("implement me")
}

func (p *StartingService) Init(merchant Merchant) {
	p.merchant = merchant
}
//
//func (p *StartingService) OnEvent(evt Event) error {
//	switch evt.GetType() {
//	case EventTypeNewSession:
//		session := &EventOfNewSession{}
//		if err := json.Unmarshal(evt.GetData(), session); err != nil {
//			return err
//		}
//
//		// 新服务开始时的回调
//		p.Callback.StartingServiceOnNewSession(evt)
//
//		// 获取商户的配置
//		timeout, err := p.Callback.StartingServiceGetTimeoutOfVisitorDoesNotAsk(evt)
//		if err != nil {
//			return err
//		}
//
//		// 新建倒计时
//		// 访客长时间无提问
//
//		ct := countdown_context.New()
//		ct.SetTimeoutCallback(timeout, func(counter uint64, args ...interface{}) {
//			session := args[0].(*EventOfNewSession)
//			p.Callback.TimeoutOfVisitorDoesNotAskOnStartingService(counter, session)
//
//			// 访客长时间无提问，机器人自动回复后，进入自动结束流程
//			SetSessionState(session.SessionId, SessionStateStopping)
//			forwardEvent(session)
//		}, session)
//
//		// 启动倒计时
//		ct.Enable()
//
//		// 保存倒计时
//		p.countdownForVisitorDoesNotAsk[evt.GetSessionId()] = ct
//		break
//	case EventTypeVisitorMessage:
//		p.countdownForVisitorDoesNotAsk[evt.GetSessionId()].Disable()
//		SetSessionState(evt.GetSessionId(), SessionStateRobotService)
//		ForwardEvent(evt)
//		break
//	default:
//		err := fmt.Errorf("开场白阶段，没有事件处理:%d", EventTypeNewSession)
//		return err
//	}
//
//	return nil
//}
