package heartbeat

/*
	心跳
*/
type Heartbeat interface {
	Init() error

	// 心跳记录
	// 每次调用此方法，重新开始计算超时
	Heartbeat(uuid string) error

	// 检查是否存在
	IsAlive(uuid string) (bool, error)
}
