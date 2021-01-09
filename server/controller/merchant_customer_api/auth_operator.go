package merchant_customer_api

import (
	"api/cs/meta"
	libs_http "api/libs/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

/*
	operator数据结构
*/
type Operator struct {
	// 会话ID
	SessionId string

	*meta.Client

	// 登陆时间
	LoginTime time.Time
	// 上下文ID
	ContextId string
}

func (c *Operator) SetContextId(uuid string) {
	c.ContextId = uuid
}
func (c *Operator) GetContextId() string {
	return c.ContextId
}

func (c *Operator) GetTenantId() uint64 {
	return 0
}

func (c *Operator) GetUserId() uint64 {
	return c.UserId
}
func (c *Operator) GetUserType() uint64 {
	return uint64(c.ClientType)
}
func (c *Operator) GetNickname() string {
	return c.Nickname
}
func (c *Operator) GetThumb() string {
	return c.Thumb
}

func (c *Operator) GetId() uint64 {
	return c.UserId
}
func (c *Operator) GetUsername() string {
	return c.Username
}

func (c *Operator) GetAuthorityId() string {
	return c.Username
}

/*
	客户的UUID
	网关通过client uuid路由消息
*/
func (c *Operator) GetUUID() string {
	return fmt.Sprintf("%s:%d:%d", c.TenantCode, c.UserType, c.UserId)
}

/*
	获取操作者信息
	@method GET
*/
func (c *Operator) Info(ctx *gin.Context) {
	libs_http.RspData(ctx, 0, "", GetOperator(ctx))
}
