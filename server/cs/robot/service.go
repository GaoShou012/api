package robot

type Service interface {
	OnInit(robot *Robot, callback *Callback) error
	OnEntry(evt Event)
	OnEvent(evt Event)
	OnExit(evt Event)
	OnClean(sessionId string)
}
