package countdown

import "time"

type Event interface {
	GetCounter() uint64
	GetTimeout() time.Duration
	GetBeginTime() *time.Time
	GetEndTime() *time.Time
	GetParams() []interface{}
}
