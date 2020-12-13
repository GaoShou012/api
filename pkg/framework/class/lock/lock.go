package lock

import "time"

type Lock interface {
	Init() error
	Lock(key string, val string, timeout time.Duration) (bool, error)
	Unlock(key string, val string) (bool, error)
}
