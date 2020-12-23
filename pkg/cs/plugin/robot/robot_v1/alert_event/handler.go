package alert_event

import (
	"cs/class/robot"
	"cs/env"
	"encoding/json"
	"fmt"
	"framework/class/broker"
	"sync"
	"time"
)

const (
	brokerTopic = "robot:alert_event"
)

type Handler struct {
	debug    bool
	mutex    sync.Mutex
	sessions map[string]*session
	broker.Broker
	callback *Callback
}

func (h *Handler) Init(b broker.Broker, callback *Callback) error {
	h.sessions = make(map[string]*session)
	h.Broker = b
	h.callback = callback
	return nil
}

func (h *Handler) Debug(enable bool) {
	h.debug = enable
}

func (h *Handler) GetSession(sessionId string) *session {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	s, ok := h.sessions[sessionId]
	if !ok {
		s = newSession(h.callback)
		h.sessions[sessionId] = s
	}

	return s
}

func (h *Handler) Run() {
	h.Broker.Subscribe(brokerTopic, func(evt broker.Event) error {
		e := &event{}

		if err := json.Unmarshal(evt.Message(), e); err != nil {
			info := fmt.Sprintf("message=%v,err=%v", evt.Message(), err)
			env.Logger.Error(info)
			return err
		}

		if env.Logger.IsDebug() {
			info := fmt.Sprintf("收到事件:会话ID=%s 类型=%s 超时=%d 消息=%s", e.SessionId, e.Type, e.Timeout, string(e.Event))
			env.Logger.Info(info)
		}

		theSession := h.GetSession(e.SessionId)

		switch e.Type {
		case robot.EventTypeEstablishSession:
			// 是否开启
			if h.callback.IsEnableOnCustomerDoesNotAsk(e.Event, e.Options) == false {
				return nil
			}

			// 是否已经处理过
			ok, err := env.Session.SetFlagNX(e.SessionId, "CustomerDoesNotAck", time.Now().String())
			if err != nil {
				return err
			}
			if !ok {
				break
			}

			// 开启倒计时
			theSession.EnableCountdownForCustomerDoesNotAsk(e.Timeout, e)
			break
		case robot.EventTypeCustomerTalk:
			// 关闭访客无提问倒计时
			theSession.DisableCountdownForCustomerDoesNotAsk()

			// 查询是否开启客服自动应答
			if h.callback.IsEnableOnCustomerServerDoesNotAnswer(e.Event, e.Options) == false {
				break
			}

			// 开启客服自动应答倒计时s
			theSession.EnableCountdownForCsDoesNotAnswer(e.Timeout, e)
			break
		case robot.EventTypeCsTalk:
			// 关闭客服自动应答倒计时
			theSession.DisableCountdownForCsDoesNotAnswer()
			break
		}

		return nil
	})
}

func (h *Handler) PublishEvent(evt robot.AlertEvent, timeout time.Duration, options map[string]string) error {
	e := &event{
		SessionId: evt.GetSessionId(),
		Type:      evt.GetEventType(),
		Timeout:   timeout,
		Event:     nil,
		Options:   options,
	}

	// debug
	if env.Logger.IsDebug() {
		info := fmt.Sprintf("会话ID:%s,事件类型:%s,超时时间:%d,选项:%v", e.SessionId, e.Type, e.Timeout, e.Options)
		env.Logger.Info("机器人AlertEvent,", info)
	}

	// encode event
	{
		j, err := json.Marshal(evt)
		if err != nil {
			return err
		}
		e.Event = j
	}

	{
		j, err := json.Marshal(e)
		if err != nil {
			return err
		}
		if err := h.Broker.Publish(brokerTopic, j); err != nil {
			return err
		}
	}

	return nil
}
