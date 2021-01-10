package models

import "time"

type IpWhitelist struct {
	Model
	MerchantId    *uint
	Ip            *string
	Desc          *string
	IpGroup       *string
	LastLoginTime *time.Time
}

func (m *IpWhitelist) GetTableName() string {
	return "ip_whitelist"
}
