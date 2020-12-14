package cs

import (
	"cs/env"
	"cs/meta"
)

/*
	加入会话
*/
func JoinSession(session meta.Session, client meta.Client) error {
	if err := env.Session.SetClient(session, client); err != nil {
		return err
	}

	if err := env.Client.SetSession(client, session.GetSessionId(), session); err != nil {
		return err
	}

	return nil
}

/*
	离开会话
*/
func LeaveSession(session meta.Session, client meta.Client) error {
	if err := env.Session.DelClient(session, client); err != nil {
		return err
	}

	if err := env.Client.DelSession(client, session.GetSessionId()); err != nil {
		return err
	}

	return nil
}
