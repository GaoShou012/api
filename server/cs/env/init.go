package env

import (
	"api/cs"
	"api/cs/class/client_event"
	"api/cs/class/session"
	"im/class/im"
)

var PersonalEvent *cs.PersonalEvent
var Session session.Session
var ClientEvent client_event.ClientEvent
var IM im.IM