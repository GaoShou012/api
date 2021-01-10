package visitors

import (
	"api/global"
	"api/models"
	"time"
)

//访客信息表
type Visitors struct {
	models.Model
	MerchantId *uint64
	Username   *string
	Nickname   *string
	//访客等级
	Level *int
	//访客标签
	Tags           *string
	Gender         *int
	Phone          *uint64
	Email          *string
	Wechat         *string
	WechatNickname *string
	QQ             *string
	QQNickname     *string
	Desc           *string //备注
	//访客Ip
	Ip         *string
	IpLocation *string
	//最后接入时间
	LastVisitTime *time.Time
}

func (m *Visitors) GetTableName() string {
	return "visitors"
}

//返回数据的总数
func (m *Visitors) Count(where string) (int64, error) {
	var count int64 = 0
	res := global.DBSlave.Table(m.GetTableName()).Where(where).Count(&count)
	if res.Error != nil {
		return count, res.Error
	}
	return count, nil
}

//func (m *Visitors) SelectByUsernameMerchantId(username string, merchantId uint64) Visitors {
//
//}

//func (m *Visitors) UpdateOnConnect(Username string, merchantId uint64, Nickname string, Ip string, level int,lastVisitTime time.Time) error {
//
//	return nil
//}
