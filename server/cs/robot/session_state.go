package robot

var sessionState map[string]SessionState

type SessionState int8

const (
	SessionStateStartingService SessionState = iota
	SessionStateRobotService
	SessionStateHumanService
	SessionStateRating
	SessionStateStopping
)

func SetSessionState(sessionId string, state SessionState) {
	sessionState[sessionId] = state
}
func GetSessionState(sessionId string) (SessionState, bool) {
	state, ok := sessionState[sessionId]
	return state, ok
}
