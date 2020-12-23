package robot

import (
	"cs/meta"
	"time"
)

/*
	条件满足，触发警告消息
	回调返回消息字节流，由机器人推送给客户端
*/
//type OnQueueWarn func(session meta.Session, client meta.Client) []byte

type Robot interface {
	Init() error

	/*
		当访客说话时
		如果会话依然在排队中，触发一次排队警告
	*/
	QueueNotification(session meta.Session, client meta.Client) error

	/*
		提醒事件
	*/
	AlertEvent(event AlertEvent, timeout time.Duration, options map[string]string) error

	/*
		启动提醒事件处理者
	*/
	AlertEventHandler() error
}
