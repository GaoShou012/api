package initialize

import (
	"api/cs"
	"api/global"
	broker_redis_pubsub "framework/plugin/broker/redis_pubsub"
	stream_redis_stream "framework/plugin/stream/redis_stream"
	channel_v1 "im/plugin/channel/v1"
	client_v1 "im/plugin/client/v1"
	"im/plugin/gateway/gateway_v1"
	im_v1 "im/plugin/im/v1"
)

func InitCS() {
	stream := stream_redis_stream.New(
		stream_redis_stream.WithRedisClient(global.RedisClient),
	)

	b := broker_redis_pubsub.New(
		broker_redis_pubsub.WithRedisClient(global.RedisClient),
	)

	ch := channel_v1.New(
		channel_v1.WithRedisClient(global.RedisClient),
		channel_v1.WithStream(stream),
	)

	cli := client_v1.New(
		client_v1.WithRedisClient(global.RedisClient),
		client_v1.WithStream(stream),
	)

	gw := gateway_v1.New(
		gateway_v1.WithBroker(b),
	)

	cs.IM = im_v1.New(
		im_v1.WithRedisClient(global.RedisClient),
		im_v1.WithBroker(b),
		im_v1.WithChannel(ch),
		im_v1.WithClient(cli),
		im_v1.WithGateway(gw),
	)
	global.IM = cs.IM
}
