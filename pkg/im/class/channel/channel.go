package channel

type Channel interface {
	// 创建频道
	Create(topic string, info Info) error
	// 删除频道
	Delete(topic string) error
	// 是否存在
	Exists(topic string) (bool, error)

	// 设置频道是否开启
	SetEnable(topic string, enable bool) error
	GetEnable(topic string) bool

	SetInfo(topic string, info Info) error
	GetInfo(topic string, info Info) error

	// 频道下的客户端列表
	Clients(topic string) (Clients, error)

	// 推送消息
	Push(topic string, message []byte) (messageId string, err error)

	// 拉取消息，正向
	Pull(topic string, lastMessageId string, count uint64) ([]Event, error)

	// 拉取消息，反向
	RevPull(topic string, lastMessageId string, count uint64) ([]Event, error)

	// 拉取消息，指定ID
	PullById(topic string, messageId string) ([]byte, error)

	// 订阅
	Subscribe(topic string, clientUUID string) error

	// 取消订阅
	UnSubscribe(topic string, clientUUID string) error
}
