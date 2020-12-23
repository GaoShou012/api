package alert_event

import (
	"context"
	"sync"
	"time"
)

type Ontimeout func(counter uint64, args ...interface{})

type Countdown struct {
	mutex     sync.Mutex
	enable    bool
	cancel    context.CancelFunc
	counter   uint64
	onTimeout Ontimeout
}

func (c *Countdown) Init() error {
	return nil
}

func (c *Countdown) OnTimeout(ontimeout Ontimeout) {
	c.onTimeout = ontimeout
}

func (c *Countdown) Enable(timeout time.Duration, args ...interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.enable {
		return
	}
	c.enable = true

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		c.cancel = cancel
		defer func() {
			c.enable = false
		}()
		select {
		case <-ctx.Done():
			return
		case <-time.After(timeout):
			c.counter++
			c.onTimeout(c.counter, args)
		}
	}()
}

func (c *Countdown) Disable() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.enable == false {
		return
	}
	c.enable = false
	c.cancel()
}
