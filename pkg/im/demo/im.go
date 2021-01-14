package main

import (
	"fmt"
	broker_redis_pubsub "framework/plugin/broker/redis_pubsub"
	stream_redis_stream "framework/plugin/stream/redis_stream"
	"framework/utils"
	"im/env"
	channel_v1 "im/plugin/channel/v1"
	client_v1 "im/plugin/client/v1"
	"im/plugin/gateway/gateway_v1"
	im_v1 "im/plugin/im/v1"
	"time"
)

type session struct {
	Id        string
	Enable    bool
	CreatedAt time.Time
}

func (s *session) GetEnable() bool {
	return s.Enable
}
func (s *session) SetEnable(enable bool) {
	s.Enable = enable
}
func (s *session) GetTopic() string {
	return fmt.Sprintf("session:%s", s.Id)
}

func main() {
	dns := fmt.Sprintf("redis://:@192.168.0.2:17001?Db=0&PoolMax=100&PoolMin=10")
	redisClient, err := utils.RedisClient(dns)
	if err != nil {
		env.Logger.Error(err)
		return
	}

	st := stream_redis_stream.New(
		stream_redis_stream.WithRedisClient(redisClient),
	)

	b := broker_redis_pubsub.New(
		broker_redis_pubsub.WithRedisClient(redisClient),
	)

	gw := gateway_v1.New(
		gateway_v1.WithBroker(b),
	)

	ch := channel_v1.New(
		channel_v1.WithRedisClient(redisClient),
		channel_v1.WithStream(st),
	)

	cli := client_v1.New(
		client_v1.WithRedisClient(redisClient),
		client_v1.WithStream(st),
	)

	im := im_v1.New(
		im_v1.WithRedisClient(redisClient),
		im_v1.WithGateway(gw),
		im_v1.WithChannel(ch),
		im_v1.WithClient(cli),
	)

	//sessionId := uuid.NewV1().String()
	sessionId := "69dbec34-51b5-11eb-b881-acde48001122"
	//sessionId := "1234"
	fmt.Println("sessionId:", sessionId)

	se := &session{
		Id:        sessionId,
		CreatedAt: time.Now(),
	}
	if err := im.CreateChannel(se.GetTopic(), se); err != nil {
		env.Logger.Error(err)
	}

	// 设置频道启用状态
	if err := im.Channel().SetEnable(se.GetTopic(), true); err != nil {
		env.Logger.Error(err)
		return
	}

	// 客户端订阅频道
	if err := im.ClientUnSubscribeChannel("bob:1:100", se.GetTopic()); err != nil {
		env.Logger.Error(err)
		return
	}

	// 客户端取消订阅频道
	if err := im.ClientSubscribeChannel("bob:1:100", se.GetTopic()); err != nil {
		env.Logger.Error(err)
		return
	}

	fmt.Println("推送消息到频道")
	msg := fmt.Sprintf("i am msg %v", time.Now().Format(time.RFC3339))
	fmt.Println(msg)
	msgId, err := im.PushMessageToChannel(se.GetTopic(), []byte(msg))
	if err != nil {
		env.Logger.Error(err)
		return
	}
	fmt.Println("msgId:", msgId)

	{
		fmt.Println("拉取频道消息")
		events, err := im.Channel().Pull(se.GetTopic(), "0", 10)
		if err != nil {
			env.Logger.Error(err)
			return
		}
		for _, event := range events {
			fmt.Println(event.Id(), string(event.Data()))
		}

		fmt.Println("拉取频道消息ByID·")
		event, err := im.Channel().PullById(se.GetTopic(), "1610374556776-0")
		if err != nil {
			env.Logger.Error(err)
			return
		}
		fmt.Println(string(event))
	}

	fmt.Println("客户端拉取消息")
	events, err := cli.Pull("bob:1:100", "0", 10)
	if err != nil {
		env.Logger.Error(err)
		return
	}
	for _, event := range events {
		fmt.Println(event.Id(), string(event.Data()))
	}

	//cli.Delete("bob:1:100")
	//return
	{
		fmt.Println("拉取客户端消息")
		events, err := im.PullMessageFromClient("bob:1:100", "0", 10)
		if err != nil {
			env.Logger.Error(err)
			return
		}
		for _, event := range events {
			fmt.Println(event.Id(), string(event.Data()))
		}
	}

	//im.Client().Delete("bob:1:100")
}
