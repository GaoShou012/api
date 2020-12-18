package env

import (
	"cs/class/client"
	"cs/class/queue"
	"cs/class/session"
)

var (
	Client  client.Client
	Session session.Session
	Queue   queue.Queue
)
