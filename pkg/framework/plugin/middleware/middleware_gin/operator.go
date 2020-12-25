package middleware_gin

type operatorTest struct {
	UserId    uint64
	Username  string
	ContextId string
}

func (c *operatorTest) SetContextId(uuid string) {
	c.ContextId = uuid
}
func (c *operatorTest) GetContextId() string {
	return c.ContextId
}
