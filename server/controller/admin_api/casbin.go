package admin_api

import (
	libs_http "api/libs/http"
	"api/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Casbin struct{}

// 拦截器
func (c *Casbin) CasbinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		operator, _ := GetOperator(ctx)
		AuthorityId := operator.Username
		// 获取请求的URI
		obj := ctx.Request.URL.RequestURI()
		// 获取请求方法
		act := ctx.Request.Method
		// 获取用户的角色
		sub := AuthorityId
		fmt.Println(sub,obj,act)

		casbinModel := models.CasbinRule{}
		success,err := casbinModel.ExecutePermission(sub, obj, act)
		if success {
			ctx.Next()
		} else {
			libs_http.RspState(ctx, 1, err)
			ctx.Abort()
			return
		}
	}
}
