package queue_event

import "cs/meta"

type Callback struct {
	Notification
}

/*
	排队通知
*/
type Notification func(session meta.Session, client meta.Client) []byte
