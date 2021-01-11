package models

import "time"

type IpBlacklist struct {
	Model
	MerchantId    *uint      //商户id
	Ip            *string    //黑名单ip
	Desc          *string    //描述
	IpGroup       *string    //ip组
	LastLoginTime *time.Time //上一次登陆时间
}

func (m *IpBlacklist) GetTableName() string {
	return "ip_whitelist"
}
