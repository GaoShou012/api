package event

import "time"

type ClientMessage struct {
	Sender      interface{}
	Content     interface{}
	ContentType string
	Time        time.Time
}

func (c *ClientMessage) Type() string {
	return "client.message"
}

func NewClientMessage()
