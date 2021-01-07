package cs

import (
	"api/cs/message"
	"time"
)

func NewMessageWithContent(sender Sender, content message.Content) *message.Message {
	client := &Client{
		TenantCode: sender.GetTenantCode(),
		UserId:     sender.GetUserId(),
		UserType:   sender.GetUserType(),
		Nickname:   sender.GetNickname(),
		Thumb:      sender.GetThumb(),
	}
	return &message.Message{
		Type:        content.GetMessageType(),
		Time:        time.Now(),
		Sender:      client,
		Content:     content,
		ContentType: content.GetContentType(),
	}
}