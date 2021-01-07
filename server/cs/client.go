package cs

import "fmt"

const (
	ClientTypeSystem = iota
	ClientTypeCustomer
	ClientTypeCustomerServer
)

type Sender interface {
	GetTenantCode() string
	GetUserId() uint64
	GetUserType() uint64
	GetNickname() string
	GetThumb() string
}

type Client struct {
	TenantCode string
	UserId     uint64
	UserType   string
	Username   string
	Nickname   string
	Thumb      string
}

func (c *Client) GetUUID() string {
	return fmt.Sprintf("%s:%d:%d", c.TenantCode, c.UserType, c.UserId)
}

func NewClientFromSender(sender Sender) *Client {
	return &Client{
		TenantCode: sender.GetTenantCode(),
		UserId:     sender.GetUserId(),
		UserType:   sender.GetUserType(),
		Nickname:   sender.GetNickname(),
		Thumb:      sender.GetThumb(),
	}
}

func NewClientSystem() *Client {
	return &Client{
		UserId:   0,
		UserType: ClientTypeSystem,
		Nickname: "系统",
		Thumb:    "",
	}
}
