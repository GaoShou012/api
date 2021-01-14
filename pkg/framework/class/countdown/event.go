package countdown

import "time"

type Event interface {
	GetCounter() uint64
	GetParams() []interface{}
	GetBeginTime() time.Time
	GetEndTime() time.Time
}
