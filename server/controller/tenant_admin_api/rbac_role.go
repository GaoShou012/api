package tenant_admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"api/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RbacRole struct {
}

func (c *RbacRole) Create(ctx *gin.Context) {
	var params models.RbacRole
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	operator := GetOperator(ctx)
	{
		role := &params
		if err := global.TenantRBAC.CreateRole(operator, role); err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
	}

	libs_http.RspState(ctx, 0, "创建角色成功")
}


func (c *RbacRole) Update(ctx *gin.Context) {
	var params models.RbacRole
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	operator := GetOperator(ctx)
	{
		roleId := *params.Id
		role := &params
		if err := global.TenantRBAC.UpdateRole(operator, roleId, role); err != nil {
			libs_http.RspState(ctx,1,err)
			return
		}
	}

	libs_http.RspState(ctx,0,"更新角色成功")
}

func (c *RbacRole) Delete(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	if err := ctx.Bind(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	operator := GetOperator(ctx)

	{
		id := params.Id
		ok, err := global.TenantRBAC.DeleteRole(operator, id)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		if !ok {
			libs_http.RspState(ctx, 1, fmt.Errorf("删除失败，可能角色ID(%d)不存在", id))
		}
	}

	libs_http.RspState(ctx, 0, "删除成功")
}
