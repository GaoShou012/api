package admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"api/models"
	"github.com/gin-gonic/gin"
)

type AuthorityRolesMenus struct{}

func (c *AuthorityRolesMenus) Get(ctx *gin.Context) {
	var params struct {
		Page int
		Size int
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &models.AuthorityRolesMenus{}

	var list []models.AuthorityRolesMenus
	// 按名称查找
	count, err := model.Count("*")
	if err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	res := global.DBSlave.Order("id desc").Limit(params.Size)
	res.Offset(params.Page * params.Size).Find(&list)
	if res.Error != nil {
		libs_http.RspState(ctx, 0, res.Error)
		return
	}


	libs_http.RspSearch(ctx, 0, nil, count,list)
	return
}

func (c *AuthorityRolesMenus) Create(ctx *gin.Context) {
	// 接受参数
	var params struct {
		RoleId     uint64
		MenuGroups string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &models.AuthorityRolesMenus{
		RoleId:&params.RoleId,
		//MenuGroups:&params.MenuGroups,
	}
	if res := global.DBMaster.Create(model); res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
	}
	libs_http.RspState(ctx, 0, "创建成功")
}

func (c *AuthorityRolesMenus) Update(ctx *gin.Context) {
	// 接受参数
	var params struct {
		Id   uint64
		RoleId     uint64
		MenuId uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	//model := &models.MenuGroups{}

	model := &models.AuthorityRolesMenus{
		Id:   &params.Id,
		RoleId: &params.RoleId,
		MenuId: &params.MenuId,
	}
	if err := model.UpdateById(model); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	libs_http.RspState(ctx, 0, "更新成功")
}

func (c *AuthorityRolesMenus) Del(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &models.AuthorityRolesMenus{
		Id: &params.Id,
	}
	err := model.DeleteById(model)
	if err != nil {
		libs_http.RspState(ctx, 0, err)
		return
	}

	libs_http.RspState(ctx, 0, "删除成功")
}
