package cache

import "cs/meta"

type Cache interface {
	SetClientInfo(key string, info meta.ClientInfo) error
	GetClientInfo(key string, info meta.ClientInfo) (bool, error)
	ExistsClientInfo(key string) (bool, error)

	SetSessionInfo(key string, info meta.SessionInfo) error
	GetSessionInfo(key string, info meta.SessionInfo) (bool, error)
	ExistsSessionInfo(key string) (bool, error)
}
