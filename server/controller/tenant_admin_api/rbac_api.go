package tenant_admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"api/models"
	"github.com/gin-gonic/gin"
)

type RbacApi struct {
}

/*
	@Desc 创建API
	@Method POST
	@Developer GaoShou
*/
func (c *RbacApi) Create(ctx *gin.Context) {
	var params struct {
		Method string
		Path   string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	operator := GetOperator(ctx)
	model := &models.RbacApi{
		Model:    models.Model{},
		Method:   &params.Method,
		Path:     &params.Path,
	}

	if err := global.TenantRBAC.CreateApi(operator, model); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	libs_http.RspState(ctx, 0, "创建成功")
}

/*
	@Desc 创建API
	@Method GET
	@Developer GaoShou
*/
func (c *RbacApi) Delete(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	if err := ctx.Bind(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	operator := GetOperator(ctx)

	{
		apiId := params.Id
		ok, err := global.TenantRBAC.DeleteApi(operator, apiId)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		if !ok {
			libs_http.RspState(ctx, 1, "删除失败，可能数据不存在")
			return
		}
	}

	libs_http.RspState(ctx, 0, "删除成功")
}

func (c *RbacApi) Update(ctx *gin.Context) {
	var params models.RbacApi
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	if *params.Id == 0 {
		libs_http.RspState(ctx, 1, "无效的ID")
		return
	}

	operator := GetOperator(ctx)

	apiId := *params.Id
	api := &params
	if err := global.TenantRBAC.UpdateApi(operator, apiId, api); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	libs_http.RspState(ctx, 0, "更新成功")
}
