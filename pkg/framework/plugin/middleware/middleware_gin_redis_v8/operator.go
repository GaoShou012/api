package middleware_gin_redis_v8

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
