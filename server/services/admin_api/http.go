package admin_api

import (
	controller_admin_api "api/controller/admin_api"
	"api/meta"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpService struct{}

func (r *HttpService) Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//请求方法
		method := ctx.Request.Method

		// 允许任何源
		ctx.Header("Access-Control-Allow-Origin", "*")
		//服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		//ctx.Header("Access-Control-Allow-Headers", "Token,Content-Type")
		ctx.Header("Access-Control-Allow-Headers", "*")
		// 跨域关键设置 让浏览器可以解析
		ctx.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			ctx.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		ctx.Next() //  处理请求
	}
}

func (r *HttpService) Route(engine *gin.Engine) {
	controller_admin_api.InitOperatorContext()

	api := engine.Group(fmt.Sprintf("/admin%s", meta.ApiVersion))
	var authenticated gin.IRoutes

	// 登陆&验证
	{
		c := controller_admin_api.Auth{}
		api.POST("/login", c.Login)
		api.POST("/register", c.Register)

		authenticated = api
		authenticated.Use(controller_admin_api.OperatorContext.Parse().(gin.HandlerFunc))
		//authenticated.Use(controller_admin_api.OperatorContext.Expiration().(gin.HandlerFunc))
		authenticated.GET("/logout", c.Logout)
	}

	// 操作者
	{
		c := controller_admin_api.Operator{}
		authenticated.GET("/operator/info", c.Info)
	}

}
