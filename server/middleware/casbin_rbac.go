package middleware

import (
	"api/global"
	libs_http "api/libs/http"
	"github.com/gin-gonic/gin"
)

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		operator, err := libs_http.GetOperator(ctx)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			ctx.Abort()
			return
		}
		AuthorityId := operator.GetAuthorityId()
		// 获取请求的URI
		obj := ctx.Request.URL.RequestURI()
		// 获取请求方法
		act := ctx.Request.Method
		// 获取用户的角色
		sub := AuthorityId
		// 判断策略中是否存在
		success, err := global.CasbinEnforcer.Enforce(sub, obj, act)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			ctx.Abort()
			return
		}
		if !success {
			libs_http.RspState(ctx, 1, "权限不足")
			ctx.Abort()
			return
		}
	}
}
