package countdown_ticker

import (
	"context"
	"framework/class/countdown"
	"sync"
	"time"
)

var _ countdown.Countdown = &plugin{}

type plugin struct {
	mutex  sync.Mutex
	enable bool
	cancel context.CancelFunc

	opts *Options

	counter        uint64
	ticker         *time.Ticker
	timeout        time.Duration
	beginTime      *time.Time
	endTime        *time.Time
	callback       countdown.Callback
	callbackParams []interface{}
}

func (p *plugin) Event() countdown.Event {
	evt := &event{
		counter:   p.counter,
		params:    p.callbackParams,
		beginTime: p.beginTime,
		endTime:   p.endTime,
	}
	return evt
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) New(t time.Duration, callback countdown.Callback, args ...interface{}) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.enable {
		return
	}
	p.enable = true

	if p.ticker == nil {
		p.ticker = time.NewTicker(t)
	} else {
		p.ticker.Reset(t)
	}

	now := time.Now()
	p.callback = callback
	p.callbackParams = args
	p.beginTime = &now
	p.endTime = nil

	go func() {
		t, ok := <-p.ticker.C
		if !ok {
			return
		}
		p.ticker.Stop()
		p.counter++
		p.enable = false
		p.endTime = &t
		evt := &event{
			timeout:   p.timeout,
			counter:   p.counter,
			params:    p.callbackParams,
			beginTime: p.beginTime,
			endTime:   p.endTime,
		}
		p.callback(evt)
	}()
}

func (p *plugin) Reset() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.ticker.Reset(p.timeout)
}

func (p *plugin) Stop() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.enable == false {
		return
	}
	p.enable = false
	p.ticker.Stop()
}
