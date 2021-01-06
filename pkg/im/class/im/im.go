package im

import (
	"im/class/channel"
	"im/class/client"
	"im/class/gateway"
)

type IM interface {
	// 获取频道适配器
	Channel() channel.Channel
	// 获取客户端适配器
	Client() client.Client
	// 获取网关
	Gateway() gateway.Gateway

	// 客户端
	ClientAttach(uuid string, lastMessageId string)
	ClientDetach(uuid string)

	// 创建频道
	CreateChannel(info channel.Info) error
	// 删除频道
	DeleteChannel(topic string) error

	// 频道的客户列表
	ClientsOfChannel(topic string) (channel.Clients, error)
	// 客户的频道列表
	ChannelsOfClient(clientUUID string) (client.Channels, error)

	// 推送消息到频道
	PushMessageToChannel(topic string, message []byte) (messageId string, err error)
	PushMessageToChannelWithClients(topic string, message []byte, clients channel.Clients) (messageId string, err error)

	// 直接推送消息给客户端
	PushMessageToClient(uuid string, message []byte) error
	// 使用stream推送消息给客户端
	PushMessageToClientEvent(uuid string, message []byte) (messageId string, err error)

	// 客户端订阅频道
	ClientSubscribeChannel(uuid string, topic string) error
	// 客户端取消订阅频道
	ClientUnSubscribeChannel(uuid string, topic string) error

	// 响应消息推送
	OnPublish(callback OnPublishCallback)
}
