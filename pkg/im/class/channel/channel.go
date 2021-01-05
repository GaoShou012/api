package channel

type Channel interface {
	// 创建频道
	Create(info Info) error
	// 删除频道
	Delete(topic string) error

	// 设置频道是否开启
	SetEnable(topic string, enable bool) error
	GetEnable(topic string) bool

	SetInfo(topic string, info Info) error
	GetInfo(topic string, info Info) error

	// 频道下的客户端列表
	Clients(topic string) (Clients, error)

	// 推送消息
	Publish(topic string, message []byte) (messageId string, err error)
	// 订阅
	Subscribe(topic string, clientUUID string) error
	// 取消订阅
	UnSubscribe(topic string, clientUUID string) error

	// 释放频道
	Release(topic string) error
}
