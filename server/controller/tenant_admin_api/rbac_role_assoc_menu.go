package tenant_admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"github.com/gin-gonic/gin"
)

type RbacRoleAssocMenu struct {

}

func (c *RbacRoleAssocMenu) Create(ctx *gin.Context){
	var params struct{
		RoleId uint64
		MenuId uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx,1,err)
		return
	}

	operator := GetOperator(ctx)
	if err := global.TenantRBAC.RoleAssocMenu(operator,params.RoleId,params.MenuId); err != nil {
		libs_http.RspState(ctx,1,err)
		return
	}

	libs_http.RspState(ctx,0,"角色关联菜单成功")
}

func (c *RbacRoleAssocMenu) Delete(ctx*gin.Context) {
	var params struct{
		Id uint64
	}
	if err := ctx.Bind(&params); err != nil {
		libs_http.RspState(ctx,1,err)
		return
	}

	operator := GetOperator(ctx)
	if err := global.TenantRBAC.RoleDisassociateMenu(operator,params.Id); err != nil {
		libs_http.RspState(ctx,1,err)
		return
	}

	libs_http.RspState(ctx,0,"角色取消关联菜单")
}