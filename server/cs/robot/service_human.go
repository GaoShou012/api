package robot

var _ Service = &ServiceHuman{}

type ServiceHuman struct{}

func (s ServiceHuman) OnInit(robot *Robot, callback *Callback) error {
	return nil
}

func (s ServiceHuman) OnEntry(evt Event) {
}

func (s ServiceHuman) OnExit(evt Event) {
}

func (s ServiceHuman) OnClean(sessionId string) {
}

func (s ServiceRating) OnEvent(evt Event) {
	//switch evt.GetType() {
	//default:
	//	err := fmt.Errorf("未知的事件类型,evt=%v", evt)
	//	env.Logger.Error(err)
	//	break
	//}
}
