package robot

import (
	"cs/env"
	"fmt"
)

var _ Service = &ServiceRating{}

type ServiceRating struct{}

func (s ServiceRating) OnEntry(evt Event) {
	panic("implement me")
}

func (s ServiceRating) OnExit(evt Event) {
	panic("implement me")
}

func (s ServiceHuman) OnEvent(evt Event) {
	switch evt.GetType() {
	default:
		err := fmt.Errorf("未知的事件类型,evt=%v", evt)
		env.Logger.Error(err)
		break
	}
}
