package admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"github.com/gin-gonic/gin"
)

type RbacRoleAssocMenuGroup struct {

}

func (c *RbacRoleAssocMenuGroup) Create(ctx *gin.Context){
	var params struct{
		RoleId uint64
		MenuGroupId uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx,1,err)
		return
	}

	operator := GetOperator(ctx)
	if err := global.RBAC.RoleAssocMenuGroup(operator,params.RoleId,params.MenuGroupId); err != nil {
		libs_http.RspState(ctx,1,err)
		return
	}

	libs_http.RspState(ctx,0,"角色关联菜单组成功")
}

func (c *RbacRoleAssocMenuGroup) Delete(ctx *gin.Context) {
	var params struct{
		Id uint64
	}
	if err := ctx.Bind(&params); err != nil {
		libs_http.RspState(ctx,1,err)
		return
	}

	operator := GetOperator(ctx)
	if err := global.RBAC.RoleDisassociateMenuGroup(operator,params.Id); err != nil{
		libs_http.RspState(ctx,1,err)
		return
	}

	libs_http.RspState(ctx,0,"角色取消菜单组成功")
}
