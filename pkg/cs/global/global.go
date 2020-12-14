package global

import (
	"cs/class"
	"cs/class/cache"
	"cs/class/session"
	"cs/meta"
)

var (
	Adapter    class.Adapter
	Session    session.Session
	Client     class.Client
	ClientInfo meta.Client

	Cache cache.Cache
)
