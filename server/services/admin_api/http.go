package admin_api

import (
	controller_admin_api "api/controller/admin_api"
	"api/meta"
	"fmt"
	"github.com/gin-gonic/gin"
)

type HttpService struct{}

func (r *HttpService) Cors() gin.HandlerFunc {
	return nil
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
		api.GET("/auth_code", c.CodeImage)

		authenticated = api
		authenticated.Use(controller_admin_api.OperatorContext.Parse().(gin.HandlerFunc))
		authenticated.Use(controller_admin_api.OperatorContext.Expiration().(gin.HandlerFunc))
		authenticated.GET("/logout", c.Logout)
	}

	// 操作者
	{
		c := controller_admin_api.Operator{}
		authenticated.GET("/operator/info", c.Info)
	}

	// RBAC API
	{
		c := controller_admin_api.RbacApi{}
		authenticated.POST("/rbac/api/create", c.Create)
		authenticated.POST("/rbac/api/update", c.Update)
		authenticated.GET("/rbac/api/delete", c.Delete)
	}

	// RBAC Menu
	{
		c := controller_admin_api.RbacMenu{}
		authenticated.POST("/rbac/menu/create", c.Create)
		authenticated.POST("/rbac/menu/update", c.Update)
		authenticated.GET("/rbac/menu/delete", c.Delete)
	}

	// RBAC Role
	{
		c := controller_admin_api.RbacRole{}
		authenticated.POST("/rbac/role/create", c.Create)
		authenticated.POST("/rbac/role/update", c.Update)
		authenticated.GET("/rbac/role/select", c.Select)
	}

	//RBAC MenuGroup
	{
		c := controller_admin_api.RbacMenuGroup{}
		authenticated.POST("/rbac/menu_Group/create", c.Create)
		authenticated.POST("/rbac/menu_Group/update", c.Update)
		authenticated.GET("/rbac/menu_Group/delete", c.Delete)
	}

	// RBAC Role Assoc API
	{
		c := controller_admin_api.RbacRoleAssocApi{}
		authenticated.POST("/rbac/role_assoc_api/create", c.Create)
		authenticated.GET("/rbac/role_assoc_api/delete", c.Delete)
	}

	// RBAC Role Assoc MenuGroup
	{
		c := controller_admin_api.RbacRoleAssocMenuGroup{}
		authenticated.POST("/rbac/role_assoc_menu_group/create", c.Create)
		authenticated.GET("/rbac/role_assoc_menu_group/delete", c.Delete)
	}
	// RBAC Role Assoc Menu
	{
		c := controller_admin_api.RbacRoleAssocMenu{}
		authenticated.POST("/rbac/role_assoc_menu/create", c.Create)
		authenticated.GET("/rbac/role_assoc_menu/delete", c.Delete)
	}
}
