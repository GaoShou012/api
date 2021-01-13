package models

import "time"

type Sessions struct {
	Model
	MerchantId         *uint64 //商户id
	Status             *int    //状态
	SessionId          *string //会话id
	UserId             *uint64
	UserType           *int       //用户类型
	UserSource         *string    //用户来源
	UserVipLevel       *int       //用户等级
	UserName           *string    //用户名
	UserIp             *string    //用户ip
	UserTerminal       *int       //用户终端
	UserToken          *string    //用户token、
	UserIpLocation     *string    //用户ip地址所在地
	UserRating         *int64     //用户评分
	UserComment        *string    //用户评语
	UserRatingTime     *time.Time //访客评价时间
	CsId               *uint64    //客服id
	CsName             *string    //客服名
	CsGroup            *int       //客服组
	CsDepartment       *int       //客服部门
	ServiceTags        *string    //服务标签
	ServiceType        *string    //服务类型
	ServiceTopic       *string    //服务主题
	CsValue            *int       //客服评估:1待定评价，2无价值，3有价值，4很有价值，5价值待定
	CsComment          *string    //客服评价备注
	ServiceRequestTime *time.Time //访客请求服务的时间,等于创建服务时间
	ServiceBeginTime   *time.Time //服务开始时间
	ServiceEndTime     *time.Time //服务结束时间
	ServiceEndReason   *string    //服务结束原因
}

func (m *Sessions) GetTableName() string {
	return "sessions"
}
