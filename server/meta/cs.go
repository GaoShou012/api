package meta

type Client struct {
	TenantCode string
	UserId     uint64
	Username   string
	UserType   string
}

func (c *Client) GetUUID() string {
	return ""
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