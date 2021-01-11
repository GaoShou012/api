package robot

import (
	"framework/class/countdown"
	"time"
)

type Callback struct {
	NewCountdownForVisitorDoesNotAskOfStartingService

	// 新会话开始时回调
	NewStartingService
	// 新会话，访客无提问回调
	TimeoutOfVisitorDoesNotAskOnStartingService
	// 获取商户在新会话无提问的超时时间
	GetTimeoutOfVisitorDoesNotAskOnStartingService
	GetTimeoutStoppingService
}

// 新建一个倒计时用于访客开始服务阶段无提问
type NewCountdownForVisitorDoesNotAskOfStartingService func(args ...interface{}) countdown.Countdown

type NewStartingService func(evt *EventOfNewSession)
type GetTimeoutOfVisitorDoesNotAskOnStartingService func(merchantCode string) (time.Duration, error)
type TimeoutOfVisitorDoesNotAskOnStartingService func(counter uint64, evt *EventOfNewSession)

// 进入自动结束流程的倒计时时间
type GetTimeoutOf
