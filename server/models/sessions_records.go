package models

import "time"

type SessionsRecords struct {
	Model
	SessionId   *uint64    //会话主键ID
	SenderId    *uint64    //发送者ID
	SenderType  *string    //发送者类型
	SenderName  *string    //发送者名称
	Message     *string    //消息内容
	MessageType *string    //消息类型
	MessageTime *time.Time //消息的发送时间
}

func (m *SessionsRecords) GetTableName() string {
	return "session_records"
}
