package admin_api

import (
	libs_http "api/libs/http"
	"github.com/gin-gonic/gin"
	"time"
)

/*
	operator数据结构
*/
type Operator struct {
	// 租户ID
	TenantId uint64
	// 租户编码
	TenantCode string
	// 用户ID
	UserId uint64
	// 用户类型
	UserType uint64
	// 用户账号
	Username string
	// 用户昵称
	Nickname string
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
	获取操作者信息
	@method GET
*/
func (c *Operator) Info(ctx *gin.Context) {
	libs_http.RspData(ctx, 0, "",GetOperator(ctx))
}
