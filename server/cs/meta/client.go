package meta

import "fmt"

const (
	ClientTypeSystem ClientType = iota
	ClientTypeCustomer
	ClientTypeCustomerServer
)

type ClientType uint64

type Client struct {
	MerchantId   uint64
	MerchantCode string
	UserType     uint64
	UserId       uint64
	Username     string
	Nickname     string
	Thumb        string
}

func (c *Client) GetUUID() string {
	return fmt.Sprintf("%s:%d:%d", c.MerchantCode, c.UserType, c.UserId)
}

func (c *Client) GetUserId() uint64 {
	return c.UserId
}
func (c *Client) GetUserType() uint64 {
	return c.UserType
}
func (c *Client) GetUsername() string {
	return c.Username
}
func (c *Client) GetNickname() string {
	return c.Nickname
}
func (c *Client) GetThumb() string {
	return c.Thumb
}
