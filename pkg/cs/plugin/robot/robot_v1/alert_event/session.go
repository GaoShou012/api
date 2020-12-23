package alert_event

import (
	lib_countdown "kfserver/libs/countdown"
	"time"
)

func newSession(callback *Callback) *session {
	s := &session{callback: callback}
	s.Init()
	return s
}

type session struct {
	countdownForCustomerDoesNotAsk *lib_countdown.Countdown
	countdownForCsDoesNotAnswer    *lib_countdown.Countdown
	callback                       *Callback
}

func (s *session) Init() {
	s.countdownForCustomerDoesNotAsk = lib_countdown.New(func(counter uint64, args ...interface{}) {
		evt, _ := args[0].([]byte)
		options, _ := args[1].(map[string]string)
		s.callback.OnCustomerDoesNotAsk(evt, options, s.countdownForCustomerDoesNotAsk.Counter())
	})
	s.countdownForCsDoesNotAnswer = lib_countdown.New(func(counter uint64, args ...interface{}) {
		evt, _ := args[0].([]byte)
		options, _ := args[1].(map[string]string)
		s.callback.OnCustomerServerDoesNotAnswer(evt, options, s.countdownForCsDoesNotAnswer.Counter())
	})
}

/*
	开启倒计时
	计数访客超时，没有作出提问
*/
func (s *session) EnableCountdownForCustomerDoesNotAsk(timeout time.Duration, evt *event) {
	s.countdownForCustomerDoesNotAsk.Enable(timeout, evt.Event, evt.Options)
}

/*
	关闭倒计时
*/
func (s *session) DisableCountdownForCustomerDoesNotAsk() {
	s.countdownForCustomerDoesNotAsk.Disable()
}

/*
	开启客服无应答倒计时
*/
func (s *session) EnableCountdownForCsDoesNotAnswer(timeout time.Duration, evt *event) {
	s.countdownForCsDoesNotAnswer.Enable(timeout, evt.Event, evt.Options)
}
func (s *session) DisableCountdownForCsDoesNotAnswer() {
	s.countdownForCsDoesNotAnswer.Disable()
}
