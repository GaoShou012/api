package notification

import (
	"api/cs/event"
)

type Notification interface {
	GetType() string
}

func NewMessage(sender interface{}, notification Notification) *event.Message {
	return event.New(event.MsgTypeNotification, sender, notification, notification.GetType())
}

