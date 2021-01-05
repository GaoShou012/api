package im

import (
	"im/class/channel"
	"im/class/client"
	"im/class/gateway"
)

type IM interface {
	// 获取频道适配器
	GetChannelAdapter() channel.Channel
	// 获取客户端适配器
	GetClientAdapter() client.Client
	GetClient() client.Client
	// 获取网关
	GetGateway() gateway.Gateway

	// 创建频道
	CreateChannel(info channel.Info) error
	// 删除频道
	DeleteChannel(topic string) error

	// 频道的客户列表
	ClientsOfChannel(topic string) (channel.Clients, error)
	// 客户的频道列表
	ChannelsOfClient(clientUUID string) (client.Channels, error)

	// 推送消息
	Publish(topic string, message []byte) (messageId string, err error)
	// 推送消息
	PublishWithClients(topic string, message []byte, clients channel.Clients) (messageId string, err error)
	// 客户端订阅频道
	Subscribe(topic string, clientUUID string) error
	// 客户端取消订阅频道
	UnSubscribe(topic string, clientUUID string) error

	RunGateway(handler GatewayHandler)
}
