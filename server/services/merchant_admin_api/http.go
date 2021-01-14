package merchant_admin_api

import (
	controller "api/controller/merchant_admin_api"
	"api/meta"
	"fmt"
	"github.com/gin-gonic/gin"
)

type HttpService struct{}

func (r *HttpService) Cors() gin.HandlerFunc {
	return nil
}

func (r *HttpService) Route(engine *gin.Engine) {
	controller.InitOperatorContext()

	api := engine.Group(fmt.Sprintf("/merchant_admin%s", meta.ApiVersion))
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
