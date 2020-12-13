package libs_trigger_timeout

import (
	"encoding/json"
	"sync"
)

type OnTimeoutHandler func(evt *OnTimeoutEvent) error

type Handler struct {
	events    map[string]*iEvent
	onTimeout OnTimeoutHandler
}

func (h *Handler) Init(onTimeout OnTimeoutHandler) {
	h.onTimeout = onTimeout
	h.events = make(map[string]*iEvent)
}

func (h *Handler) OnEvent(evt Event) error {
	// 创建事件
	thisEvent, ok := h.events[evt.GetUUID()]
	if !ok {
		j, err := json.Marshal(evt)
		if err != nil {
			return err
		}

		thisEvent = &iEvent{
			uuid:           evt.GetUUID(),
			state:          0,
			mutex:          sync.Mutex{},
			timeout:        evt.GetTimeout(),
			timeoutCounter: 0,
			cancel:         nil,
			onTimeout:      h.onTimeout,
			info:           j,
			infoVersion:    evt.GetInfoVersion(),
		}
		h.events[evt.GetUUID()] = thisEvent
	}

	// 执行倒计时控制
	if evt.IsCountdownEnable() {
		go thisEvent.EnableCountdown()
	} else {
		go thisEvent.DisableCountdown()
	}

	return nil
}
