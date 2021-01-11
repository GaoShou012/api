package robot

import "time"

type Merchant interface {
	GetTimeoutOfVisitorDoesNotAskOnStartingService(merchantCode string) (time.Duration,error)
}
