package sys_v1

import (
	"cs/class/sys"
	"cs/env"
	"cs/meta"
)

var _ sys.Sys = &plugin{}

type plugin struct {
	opts *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) CreateSession(tenantCode string, client meta.Client, session meta.Session) error {
	return env.Session.Create(session, client)
}

func (p *plugin) DeleteSession(session meta.Session) error {
	return env.Session.Delete(session)
}

func (p *plugin) JoinSession(client meta.Client, session meta.Session) error {
	// 客户关联会话
	if err := env.Client.SetSession(client, session); err != nil {
		return err
	}
	// 会话关联客户
	if err := env.Session.SetClient(session, client); err != nil {
		return err
	}
	return nil
}

func (p *plugin) LeaveSession(client meta.Client, session meta.Session) error {
	// 客户取消关联会话
	if err := env.Client.DelSession(client, session); err != nil {
		return err
	}
	// 会话取消关联客户
	if err := env.Session.DelClient(session, client); err != nil {
		return err
	}
	return nil
}

func (p *plugin) IsClientInSession(client meta.Client, sessionId string) (bool, error) {
	return env.Client.ExistsSession(client, sessionId)
}

func (p *plugin) Broadcast(session meta.Session, data []byte) error {
	clients, err := env.Session.GetAllClients(session)
	if err != nil {
		return err
	}
	for _, client := range clients {
		if err := env.Gateway.Publish(client.GetUUID(), data); err != nil {
			return err
		}
	}
	return nil
}
