package robot

type SessionStage int8

const (
	SessionStageStarting SessionStage = iota
	SessionStageRobotServicing
	SessionStageHumanServicing
	SessionStageRating
	SessionStageStopping
)
