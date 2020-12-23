package alert_event

import (
	"cs/class/robot"
	"time"
)

type event struct {
	SessionId string
	Type      robot.EventType
	Timeout   time.Duration
	Event     []byte
	Options   map[string]string
}
