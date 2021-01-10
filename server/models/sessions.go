package models

import "time"

type Sessions struct {
	Model
	Status             *int
	MerchantId         *uint64
	SessionId          *string
	UserId             *uint64
	UserSource         *string
	UserVipLevel       *int
	UserName           *string
	UserIp             *string
	UserDevice         *int
	UserToken          *string
	UserLocation       *string
	UserRating         *int64
	UserComment        *string
	UserRatingTime     *time.Time
	CsId               *uint64
	CsName             *string
	CsGroup            *int
	CsDepartment       *int
	ServiceTags        *string
	ServiceType        *string
	ServiceTopic       *string
	CsValue            *int
	CsComment          *string
	ServiceRequestTime *time.Time
	ServiceBeginTie    *time.Time
	ServiceEndTime     *time.Time
	ServiceEndReason   *string
}

func (m *Sessions) GetTableName() string {
	return "sessions"
}
