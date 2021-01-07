package meta

import "fmt"

type Client struct {
	// 租户ID
	TenantId uint64
	// 租户编码
	TenantCode string
	// 用户类型
	UserType string
	// 用户ID
	UserId uint64
	// 用户账号
	Username string
	// 用户昵称
	Nickname string
	// 用户头像
	Thumb string
}

func (c *Client) GetUUID() string {
	return fmt.Sprintf("%s:%s:%d", c.TenantCode, c.UserType, c.UserId)
}
func (c *Client) GetUserType() string {
	return ""
}

type Session struct {
	Id string
}

func (s *Session) GetSessionId() string {
	return s.Id
}
