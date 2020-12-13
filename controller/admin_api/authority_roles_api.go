package admin_api

import (
	libs_http "api/libs/http"
	"api/models"
	"api/utils"
	"github.com/gin-gonic/gin"
)

type AuthorityRolesApi struct{}

func (c *AuthorityRolesApi) Get(ctx *gin.Context) {
	var params struct {
		Page int
		Size int
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &models.AuthorityRolesApis{}

	var list []models.AuthorityRolesApis
	data := make(map[string]interface{})

	count, err := model.Count("*")
	if err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	res := utils.IMysql.Slave.Order("id desc").Limit(params.Size)
	res.Offset(params.Page * params.Size).Find(&list)
	if res.Error != nil {
		libs_http.RspState(ctx, 0, res.Error)
		return
	}

	data["count"] = count
	data["data"] = list

	libs_http.RspData(ctx, 0, nil, data)
	return
}

func (c *AuthorityRolesApi) Create(ctx *gin.Context) {
	// 接受参数
	var params struct {
		RoleId    uint64
		ApiMethod string
		ApiPath   string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &models.AuthorityRolesApis{
		RoleId:    &params.RoleId,
		ApiMethod: &params.ApiMethod,
		ApiPath:   &params.ApiPath,
	}
	if res := utils.IMysql.Master.Create(model); res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
	}
	libs_http.RspState(ctx, 0, "创建成功")
}

func (c *AuthorityRolesApi) Update(ctx *gin.Context) {
	// 接受参数
	var params struct {
		Id        uint64
		RoleId    uint64
		ApiMethod string
		ApiPath   string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	model := &models.AuthorityRolesApis{
		Id:        &params.Id,
		RoleId:    &params.RoleId,
		ApiMethod: &params.ApiMethod,
		ApiPath:   &params.ApiPath,
	}
	if err := model.UpdateById(model); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	libs_http.RspState(ctx, 0, "更新成功")
}

func (c *AuthorityRolesApi) Del(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &models.AuthorityRolesApis{
		Id: &params.Id,
	}
	err := model.DeleteById(model)
	if err != nil {
		libs_http.RspState(ctx, 0, err)
		return
	}

	libs_http.RspState(ctx, 0, "删除成功")
}
