package robot

import (
	"framework/class/countdown"
	countdown_ticker "framework/plugin/countdown"
)

func NewCountdown() countdown.Countdown {
	return countdown_ticker.New()
}
