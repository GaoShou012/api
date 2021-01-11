package countdown

import "time"

type Countdown interface {
	Init() error
	// 当倒计时，超时的时候触发
	SetTimeoutCallback(timeout time.Duration, onTimeout OnTimeout, args ...interface{})
	// 获取倒计时触发的次数
	Counter() uint64
	// 启用倒计时
	Enable()
	// 关闭倒计时
	Disable()
}
