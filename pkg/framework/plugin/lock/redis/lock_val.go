package lock_redis_set

import "time"

type lockVal struct {
	Val  string
	Time time.Time
}
