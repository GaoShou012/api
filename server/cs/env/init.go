package env

import (
	"api/cs"
	"api/cs/class/client_event"
	"api/cs/class/session"
	"framework/class/sortedset"
	"im/class/im"
)

var PersonalEvent *cs.PersonalEvent
var Session session.Session
var ClientEvent client_event.ClientEvent
var IM im.IM

var SessionsSet sortedset.Sortedset