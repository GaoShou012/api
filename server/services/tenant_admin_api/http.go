package tenant_admin_api

import (
	controller "api/controller/merchant_admin_api"
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
	controller.InitOperatorContext()

	api := engine.Group(fmt.Sprintf("/tenant_admin%s", meta.ApiVersion))
	var authenticated gin.IRoutes

	// 登陆&验证
	{
		c := controller.Auth{}
		api.POST("/login", c.Login)
		api.GET("/verification_code", c.CodeImage)

		authenticated = api
		authenticated.Use(controller.OperatorContext.Parse().(gin.HandlerFunc))
		authenticated.Use(controller.OperatorContext.Expiration().(gin.HandlerFunc))
		authenticated.GET("/logout", c.Logout)
	}

	// 操作者
	{
		c := controller.Operator{}
		authenticated.GET("/operator/info", c.Info)
	}

	// RBAC API
	{
		c := controller.RbacApi{}
		authenticated.POST("/rbac/api/create", c.Create)
		authenticated.POST("/rbac/api/update", c.Update)
		authenticated.GET("/rbac/api/delete", c.Delete)
	}

	// RBAC Menu
	{
		c := controller.RbacMenu{}
		authenticated.POST("/rbac/menu/create", c.Create)
		authenticated.POST("/rbac/menu/update", c.Update)
		authenticated.GET("/rbac/menu/delete", c.Delete)
	}
	//RBAC MenuGroup
	{
		c := controller.RbacMenuGroup{}
		authenticated.POST("/rbac/menu_Group/create", c.Create)
		authenticated.POST("/rbac/menu_Group/update", c.Update)
		authenticated.GET("/rbac/menu_Group/delete", c.Delete)
	}
	// RBAC Role Assoc API
	{
		c := controller.RbacRoleAssocApi{}
		authenticated.POST("/rbac/role_assoc_api/create", c.Create)
		authenticated.GET("/rbac/role_assoc_api/delete", c.Delete)
	}
	// RBAC Role Assoc MenuGroup
	{
		c := controller.RbacRoleAssocMenuGroup{}
		authenticated.POST("/rbac/role_assoc_menu_group/create", c.Create)
		authenticated.GET("/rbac/role_assoc_menu_group/delete", c.Delete)
	}
	// RBAC Role Assoc Menu
	{
		c := controller.RbacRoleAssocMenu{}
		authenticated.POST("/rbac/role_assoc_menu/create", c.Create)
		authenticated.GET("/rbac/role_assoc_menu/delete", c.Delete)
	}
}
