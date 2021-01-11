package robot

type Robot interface {
	Run() error

	// 机器人触发事件
	OnEvent(event Event) error
}
