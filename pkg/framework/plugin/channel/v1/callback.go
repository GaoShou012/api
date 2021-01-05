package channel_v1

import "framework/class/channel"

type Callback struct {
	Publish
}

type Publish func(topic string, messageId string, clients []*channel.Client)
