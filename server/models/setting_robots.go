package models

//机器人设置
type SettingRobots struct {
	Model
	AutoClose    *uint //自动结束阶段 提示N秒后，自动关闭会话 （提示，自动关闭会话）
	AfterAskTime *uint //回复指定话术后，?秒依然未提问(回复指定话术后，30秒依然未提问)
	WaitAsk      *uint //等待访客提问时间 (自动应答阶段)
}

func (m *SettingRobots) GetTableName() string {
	return "setting_robots"
}
