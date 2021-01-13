package countdown

import "time"

//type Countdown interface {
//	Init() error
//	// 倒计时复位，重新开始计时
//	Reset()
//	// 当倒计时，超时的时候触发
//	SetTimeoutCallback(timeout time.Duration, onTimeout OnTimeout, args ...interface{})
//	// 获取倒计时触发的次数
//	Counter() uint64
//	// 启用倒计时
//	Enable()
//	// 关闭倒计时
//	Disable()
//}

type Countdown interface {
	Init() error

	Config(t time.Duration, fn func(event Event), args ...interface{})

	// 复位倒计时，从新开始计时
	Reset()

	// 启用倒计时
	Enable()

	// 禁用倒计时
	Disable()
}
