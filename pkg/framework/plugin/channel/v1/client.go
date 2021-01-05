package channel_v1

import "time"

type Client struct {
	UUID      string
	Timeout   time.Duration
	CreatedAt time.Time
}

type ClientNotice interface {
	Push(uuid string, v interface{}) error
	Pull(uuid string, v interface{})
}


