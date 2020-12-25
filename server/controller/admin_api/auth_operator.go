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
	UserId    uint64
	UserType  uint64
	Username  string
	Nickname  string
	LoginTime time.Time
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
	operator, err := GetOperator(ctx)
	if err != nil {
		libs_http.RspState(ctx, 1000, err)
		return
	}
	libs_http.RspData(ctx, 0, "获取成功", operator)
}
