package message

import (
	"encoding/json"
	"time"
)

type MsgType uint8

const (
	// 消息
	MsgTypeCommon MsgType = iota
	// 系统提示
	MsgTypeSystemNotice
	// 通知
	MsgTypeNotification
	// 租户自定义
	MsgTypeCustomize
)

type Content interface {
	GetMessageType() MsgType
	GetContentType() string
}

type Message struct {
	Type        MsgType
	Time        time.Time
	Sender      interface{}
	Content     interface{}
	ContentType string
}

func New(msgType MsgType, sender interface{}, content interface{}, contentType string) *Message {
	return &Message{
		Type:        msgType,
		Time:        time.Now(),
		Sender:      sender,
		Content:     content,
		ContentType: contentType,
	}
}

func Encode(message *Message) ([]byte, error) {
	return json.Marshal(message)
}
