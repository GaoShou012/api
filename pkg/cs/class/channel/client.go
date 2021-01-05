package channel

import "time"

type Client struct {
	UUID      string
	Timeout   time.Duration
	CreatedAt time.Time
}
