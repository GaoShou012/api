package robot

import (
	"cs/env"
	"fmt"
)

var _ Service = &ServiceHuman{}

type ServiceHuman struct {}

func (s ServiceHuman) OnEntry(evt Event) {
	panic("implement me")
}

func (s ServiceHuman) OnExit(evt Event) {
	panic("implement me")
}

func (s ServiceRating) OnEvent(evt Event) {
	switch evt.GetType() {
	default:
		err := fmt.Errorf("未知的事件类型,evt=%v", evt)
		env.Logger.Error(err)
		break
	}
}

