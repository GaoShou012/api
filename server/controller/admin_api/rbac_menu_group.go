package admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"api/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RbacMenuGroup struct {
}

func (c *RbacMenuGroup) Create(ctx *gin.Context) {
	var params models.RbacMenuGroup
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	operator := GetOperator(ctx)
	{

		menuGroup := &params
		if err := global.RBAC.CreateMenuGroup(operator,menuGroup); err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
	}

	libs_http.RspState(ctx, 0, "创建菜单组成功")
}

func (c *RbacMenuGroup) Update(ctx *gin.Context){
	var params models.RbacMenuGroup
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	operator := GetOperator(ctx)
	{
		menuGroupId := *params.Id
		menuGroup := &params
		if err := global.RBAC.UpdateMenuGroup(operator, menuGroupId, menuGroup); err != nil {
			libs_http.RspState(ctx,1,err)
			return
		}
	}

	libs_http.RspState(ctx,0,"更新菜单组成功")
}

func (c *RbacMenuGroup) Delete(ctx *gin.Context) {
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
		ok, err := global.RBAC.DeleteMenuGroup(operator, id)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		if !ok {
			libs_http.RspState(ctx, 1, fmt.Errorf("删除失败，可能菜单组ID(%d)不存在", id))
		}
	}

	libs_http.RspState(ctx, 0, "删除成功")
}