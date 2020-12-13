package libs_trigger_timeout

import (
	"context"
	"sync"
	"time"
)

type Event interface {
	// 获取事件唯一标示符
	GetUUID() string

	// 获取是否开启倒计时
	IsCountdownEnable() bool

	// 获取超时时间
	GetTimeout() time.Duration

	// 获取信息版本
	GetInfoVersion() string
}

type OnTimeoutEvent struct {
	TimeoutCounter uint64
	InfoVersion    string
	Info           []byte
}

const (
	EventStateIsStop = iota
	EventStateIsRunning
)

type iEvent struct {
	uuid           string
	state          int8
	mutex          sync.Mutex
	timeout        time.Duration
	timeoutCounter uint64
	cancel         context.CancelFunc
	onTimeout      OnTimeoutHandler

	info        []byte
	infoVersion string
}

func (e *iEvent) Init(uuid string, onTimeout OnTimeoutHandler) {
	e.uuid = uuid
	e.onTimeout = onTimeout
}

func (e *iEvent) EnableCountdown() {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	// 如果已经开启，直接退出，避免重复开启
	if e.state == EventStateIsRunning {
		return
	}

	// 标记开启状态
	e.state = EventStateIsRunning

	ctx, cancel := context.WithCancel(context.Background())
	e.cancel = cancel
	select {
	case <-ctx.Done():
		return
	case <-time.After(e.timeout):
		e.timeoutCounter++
		evt := &OnTimeoutEvent{
			TimeoutCounter: e.timeoutCounter,
			InfoVersion:    e.infoVersion,
			Info:           e.info,
		}
		e.onTimeout(evt)
		return
	}
}
func (e *iEvent) DisableCountdown() {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	// 如果已经关闭，直接退出，避免重复关闭
	if e.state == EventStateIsStop {
		return
	}

	// 标记关闭状态
	e.state = EventStateIsStop

	// 关闭倒计时
	e.cancel()
}
