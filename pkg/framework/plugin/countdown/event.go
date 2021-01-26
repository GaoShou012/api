package countdown_ticker

import "time"

type event struct {
	counter   uint64
	params    []interface{}
	timeout   time.Duration
	beginTime *time.Time
	endTime   *time.Time
}

func (e *event) GetCounter() uint64 {
	return e.counter
}

func (e *event) GetTimeout() time.Duration {
	return e.timeout
}

func (e *event) GetParams() []interface{} {
	return e.params
}
func (e *event) GetBeginTime() *time.Time {
	return e.beginTime
}
func (e *event) GetEndTime() *time.Time {
	return e.endTime
}
