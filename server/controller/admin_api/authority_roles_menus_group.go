package admin_api

import (
	libs_http "api/libs/http"
	"api/models"
	"api/utils"
	"github.com/gin-gonic/gin"
)

type AuthorityRolesMenusGroup struct{}

func (c *AuthorityRolesMenusGroup) Get(ctx *gin.Context) {
	var params struct {
		Page int
		Size int
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &models.AuthorityRolesMenusGroups{}

	var list []models.AuthorityRolesMenusGroups
	data := make(map[string]interface{})
	// 按名称查找
	// 按名称查找
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

func (c *AuthorityRolesMenusGroup) Create(ctx *gin.Context) {
	// 接受参数
	var params struct {
		RoleId     uint64
		MenuGroups string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &models.AuthorityRolesMenusGroups{
		RoleId:&params.RoleId,
		MenuGroups:&params.MenuGroups,
	}
	if res := utils.IMysql.Master.Create(model); res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
	}
	libs_http.RspState(ctx, 0, "创建成功")
}

func (c *AuthorityRolesMenusGroup) Update(ctx *gin.Context) {
	// 接受参数
	var params struct {
		Id   uint64
		RoleId     uint64
		MenuGroups string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	//model := &models.MenuGroups{}

	model := &models.AuthorityRolesMenusGroups{
		Id:   &params.Id,
		RoleId: &params.RoleId,
		MenuGroups: &params.MenuGroups,
	}
	if err := model.UpdateById(model); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	libs_http.RspState(ctx, 0, "更新成功")
}

func (c *AuthorityRolesMenusGroup) Del(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &models.AuthorityRolesMenusGroups{
		Id: &params.Id,
	}
	err := model.DeleteById(model)
	if err != nil {
		libs_http.RspState(ctx, 0, err)
		return
	}

	libs_http.RspState(ctx, 0, "删除成功")
}
