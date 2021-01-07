package event

type MonitorTotalSessions struct {
	Data uint64
}

func (m *MonitorTotalSessions) Type() string {
	return "monitor.total_sessions"
}
