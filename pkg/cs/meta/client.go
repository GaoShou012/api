package meta

import "time"

type ClientInfo struct {
	UUID        string
	Type        string
	InfoVersion string
	Info        []byte
	CreatedAt   time.Time
}

type Client interface {
	GetUUID() string
}
