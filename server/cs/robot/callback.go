package robot

import (
	"time"
)

type Callback struct {
	*CallbackOfStartingService
	*CallbackOfRobotService
	*CallbackOfStoppingService
}

type GetTimeout func(merchantCode string) (time.Duration, error)
type Call func(evt Event)

type CallbackOfStartingService struct {
	OnEntry       Call
	S1Timeout     GetTimeout
	S2Timeout     GetTimeout
	S1TimeoutCall Call
	S2TimeoutCall Call
}

type CallbackOfRobotService struct {
	OnEntry         Call
	OnMessage       Call
	S1Timeout       GetTimeout
	S2Timeout       GetTimeout
	S1OnTimeoutCall Call
	S2OnTimeoutCall Call
}

type CallbackOfStoppingService struct {
	OnEntry       Call
	S1Timeout     GetTimeout
	S1TimeoutCall Call
}
