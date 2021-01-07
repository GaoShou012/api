package merchant_admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"github.com/gin-gonic/gin"
)

type RbacRoleAssocApi struct {
}

func (c *RbacRoleAssocApi) Create(ctx *gin.Context) {
	var params struct {
		RoleId uint64
		ApiId  uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	operator := GetOperator(ctx)
	if err := global.TenantRBAC.RoleAssocApi(operator, params.RoleId, params.ApiId); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	libs_http.RspState(ctx, 0, "角色关联API成功")
}
func (c *RbacRoleAssocApi) Delete(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	if err := ctx.Bind(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	operator := GetOperator(ctx)

	{
		assocId := params.Id
		if err := global.TenantRBAC.RoleDisassociateApi(operator, assocId); err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
	}

	libs_http.RspState(ctx, 0, "取消关联")
}
