package countdown_context

import (
	"context"
	"framework/class/countdown"
	"sync"
	"time"
)

var _ countdown.Countdown = &plugin{}

type plugin struct {
	mutex           sync.Mutex
	enable          bool
	cancel          context.CancelFunc
	counter         uint64
	timeout         time.Duration
	onTimeout       countdown.OnTimeout
	onTimeoutParams interface{}
	opts            *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) SetTimeoutCallback(timeout time.Duration, onTimeout countdown.OnTimeout, args ...interface{}) {
	p.timeout = timeout
	p.onTimeout = onTimeout
	p.onTimeoutParams = args
}

func (p *plugin) Counter() uint64 {
	return p.counter
}

func (p *plugin) Enable() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.enable {
		return
	}
	p.enable = true

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		p.cancel = cancel
		defer func() {
			p.enable = false
		}()
		select {
		case <-ctx.Done():
			return
		case <-time.After(p.timeout):
			p.counter++
			p.onTimeout(p.counter, p.onTimeoutParams)
		}
	}()
}

func (p *plugin) Disable() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.enable == false {
		return
	}
	p.enable = false
	p.cancel()
}
