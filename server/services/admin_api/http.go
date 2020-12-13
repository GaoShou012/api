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
	api := engine.Group(fmt.Sprintf("/admin/%s", meta.ApiVersion))
	var authenticated gin.IRoutes
	// 登陆
	{
		c := controller_admin_api.Auth{}
		api.POST("/login", c.Login)
		api.POST("/register", c.Register)
	}
	// 验证
	{

		c := controller_admin_api.Auth{}
		authenticated = api.Use(c.Verify)
		authenticated.GET("/logout", c.Logout)
	}
	//菜单
	{
		c := controller_admin_api.Menu{}
		authenticated.GET("/menu/get",c.Get)
		authenticated.POST("/menu/add",c.Create)
		authenticated.POST("/menu/up",c.Update)
		authenticated.POST("/menu/del",c.Del)
	}
	{
		c := controller_admin_api.MenuGroup{}
		authenticated.GET("/menu_group/get",c.Get)
		authenticated.POST("/menu_group/add",c.Create)
		authenticated.POST("/menu_group/up",c.Update)
		authenticated.POST("/menu_group/del",c.Del)
	}
	{
		c := controller_admin_api.AuthorityApi{}
		authenticated.GET("/authority_api/get",c.Get)
		authenticated.POST("/authority_api/add",c.Create)
		authenticated.POST("/authority_api/up",c.Update)
		authenticated.POST("/authority_api/del",c.Del)
	}
	{
		c := controller_admin_api.AuthorityRole{}
		authenticated.GET("/authority_role/get",c.Get)
		authenticated.POST("/authority_role/add",c.Create)
		authenticated.POST("/authority_role/up",c.Update)
		authenticated.POST("/authority_role/del",c.Del)
	}
	{
		c := controller_admin_api.AuthorityRolesApi{}
		authenticated.GET("/authority_roles_api/get",c.Get)
		authenticated.POST("/authority_roles_api/add",c.Create)
		authenticated.POST("/authority_roles_api/up",c.Update)
		authenticated.POST("/authority_roles_api/del",c.Del)
	}
	{
		c := controller_admin_api.AuthorityRolesMenusGroup{}
		authenticated.GET("/authority_roles_menus_group/get",c.Get)
		authenticated.POST("/authority_roles_menus_group/add",c.Create)
		authenticated.POST("/authority_roles_menus_group/up",c.Update)
		authenticated.POST("/authority_roles_menus_group/del",c.Del)
	}
	// 操作者
	{
		c := controller_admin_api.Operator{}
		authenticated.GET("/operator/info", c.Info)
	}
}
