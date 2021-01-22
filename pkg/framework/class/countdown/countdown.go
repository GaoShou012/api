package countdown

import "time"

type Countdown interface {
	Init() error

	// 新的倒计时开始
	// 如果先前有一个倒计时已经开始，会关闭旧的倒计时，重新开始一个新的倒计时
	New(t time.Duration, callback Callback, args ...interface{})

	// 获取事件，此事件记录了倒计时的超时次数，开始事件，结束事件，附加参数
	// 如果为结束，超时事件为nil
	Event() Event

	// 使定时器进行复位，如果定时器没有启用的情况下，将复位无效
	Reset()

	// 停止倒计时
	Stop()
}
