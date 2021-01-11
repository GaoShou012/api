package models

import "time"

type IpWhitelist struct {
	Model
	MerchantId    *uint      //商户id
	Ip            *string    //白名单ip
	Desc          *string    //描述
	IpGroup       *string    //ip组
	LastLoginTime *time.Time //上一次登陆时间
}

func (m *IpWhitelist) GetTableName() string {
	return "ip_whitelist"
}
