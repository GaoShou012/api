package notification

import "api/cs/message"

type Notification interface {
	GetType() string
}

func NewMessage(sender interface{}, notification Notification) *message.Message {
	return message.New(message.MsgTypeNotification, sender, notification, notification.GetType())
}
