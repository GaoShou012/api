package env

import (
	"cs/class/client"
	"cs/class/gateway"
	"cs/class/queue"
	"cs/class/robot"
	"cs/class/session"
	"cs/class/tenant"
	"framework/class/logger"
	env2 "framework/env"
)

var (
	Tenant tenant.Tenant
	Client  client.Client
	Session session.Session
	Queue   queue.Queue
	Gateway gateway.Gateway
	Robot   robot.Robot
	Logger  logger.Logger
)

func init() {
	Logger = env2.Logger
}
