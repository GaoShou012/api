package client

type Client interface {

	// 发送事件
	Push(uuid string, message []byte) (messageId string, err error)

	// 发送事件，到多个客户端
	PushClients(uuids []string, message []byte) error

	// 拉取事件
	Pull(uuid string, lastMessageId string, count uint64) ([]Event, error)

	// 客户端订阅频道
	Subscribe(uuid string, topic string) error

	// 客户端取消订阅频道
	UnSubscribe(uuid string, topic string) error

	// 频道列表
	Channels(uuid string) (Channels, error)
}
