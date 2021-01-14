package models

type IpWhitelist struct {
	Model
	MerchantId *uint64 //商户id
	Ip         *string //白名单ip
	Desc       *string //描述
}

func (m *IpWhitelist) GetTableName() string {
	return "ip_whitelist"
}
