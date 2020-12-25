package admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"api/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RbacMenu struct {
}

func (c *RbacMenu) Create(ctx *gin.Context) {
	var params models.RbacMenu
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	operator := GetOperator(ctx)
	{
		menuGroupId := *params.GroupId
		menu := &params
		if err := global.RBAC.CreateMenu(operator, menuGroupId, menu); err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
	}

	libs_http.RspState(ctx, 0, "创建菜单成功")
}

func (c *RbacMenu) Delete(ctx *gin.Context) {
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
		ok, err := global.RBAC.DeleteMenu(operator, id)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		if !ok {
			libs_http.RspState(ctx, 1, fmt.Errorf("删除失败，可能菜单ID(%d)不存在", id))
		}
	}

	libs_http.RspState(ctx, 0, "删除成功")
}

func (c *RbacMenu) Update(ctx *gin.Context) {
	var params models.RbacMenu
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	operator := GetOperator(ctx)
	{
		menuId := *params.Id
		menu := &params
		if err := global.RBAC.UpdateMenu(operator, menuId, menu); err != nil {
			libs_http.RspState(ctx,1,err)
			return
		}
	}

	libs_http.RspState(ctx,0,"更新菜单成功")
}