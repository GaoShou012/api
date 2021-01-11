package robot

import "sync"

var forwardEvent chan Event
var forwardEventInit sync.Once

func ForwardEvent(evt Event) {
	forwardEventInit.Do(func() {
		forwardEvent = make(chan Event)
		go func() {
			for {
				evt := <-forwardEvent
				RobotAgent.OnEvent(evt)
			}
		}()
	})
	forwardEvent <- evt
}
