package admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"github.com/gin-gonic/gin"
)

type Rbac struct {}

func (c *Rbac) Enforcer(ctx *gin.Context) {
	operator := GetOperator(ctx)

	roles := operator.Roles
	method := ctx.Request.Method
	path := ctx.Request.RequestURI
	if err := global.RBAC.Enforcer(operator, roles, method, path); err != nil {
		libs_http.RspAuthFailed(ctx, 1, err)
		ctx.Abort()
	}
}
