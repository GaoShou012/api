package robot

import (
	"cs/env"
	"encoding/json"
	"fmt"
	"framework/class/broker"
)

type Robot struct {
	broker.Broker
	Callback
	sessionStage map[string]SessionStage
	forwardEvent chan Event
	services     []Service
}

func (p *Robot) GetSessionStage(sessionId string) SessionStage {
	return p.sessionStage[sessionId]
}
func (p *Robot) SetSessionStage(sessionId string, stage SessionStage) {
	p.sessionStage[sessionId] = stage
}

func (p *Robot) OnInit() {
	p.services = append(p.services, AgentOfStartingService)
	p.services = append(p.services, AgentOfRobotServicing)
	p.services = append(p.services, AgentOfHumanServicing)
	p.services = append(p.services, AgentOfRating)
	p.services = append(p.services, AgentOfStoppingService)
}

func (p *Robot) OnClean(sessionId string) {
	for _, service := range p.services {
		service.OnClean(sessionId)
	}
}

func (p *Robot) OnEvent(evt Event) {
	switch p.GetSessionStage(evt.GetSessionId()) {
	case SessionStageStarting:
		AgentOfStartingService.OnEvent(evt)
		break
	case SessionStageRobotServicing:
		AgentOfRobotServicing.OnEvent(evt)
		break
	case SessionStageHumanServicing:
		AgentOfHumanServicing.OnEvent(evt)
		break
	case SessionStageRating:
		AgentOfRating.OnEvent(evt)
		break
	case SessionStageStopping:
		AgentOfStoppingService.OnEvent(evt)
		break
	default:
		err := fmt.Errorf("未知的会话阶段")
		env.Logger.Error(err)
		break
	}
}

func (p *Robot) Push(eventType EventType, sessionId string, merchantCode string, data []byte) error {
	robotEvent := &event{
		T: eventType,
		D: data,
		S: sessionId,
		M: merchantCode,
	}

	{
		j, err := json.Marshal(robotEvent)
		if err != nil {
			return err
		}
		return p.Broker.Publish("Robot", j)
	}
}

func (p *Robot) Handler() {
	p.Broker.Subscribe("Robot", func(evt broker.Event) error {
		defer evt.Ack()
		robotEvent := &event{}
		if err := json.Unmarshal(evt.Message(), robotEvent); err != nil {
			return err
		}
		p.OnEvent(robotEvent)
		return nil
	})
}
