package models

import "time"

type SessionsRecords struct {
	Model
	SessionId   *uint64
	SenderId    *uint64
	SenderType  *string
	SenderName  *string
	Message     *string
	MessageType *string
	MessageTime *time.Time
}

func (m *SessionsRecords) GetTableName() string {
	return "session_records"
}
