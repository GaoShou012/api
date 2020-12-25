package admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"api/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Menu struct{}

func (c *Menu) Get(ctx *gin.Context) {
	var params struct {
		Name string
		Page int
		Size int
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	menus := &models.Menus{}

	var list []models.Menus
	// 按名称查找
	if params.Name != "" {
		filed := fmt.Sprintf("`name` LIKE '%%%s%%'", params.Name)
		count, err := menus.Count(filed)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		res := global.DBSlave.Order("id desc").Limit(params.Size)
		res.Offset(params.Page * params.Size).Where(filed).Find(&list)
		if res.Error != nil {
			libs_http.RspState(ctx, 0, res.Error)
			return
		}
		libs_http.RspSearch(ctx, 0, nil, count,list)
		return
	}
	// 按名称查找
	count, err := menus.Count("*")
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

func (c *Menu) Create(ctx *gin.Context) {
	// 接受参数
	var params struct {
		Name    string
		GroupId uint64
		Sort    uint64
		Icon    string
		Path    string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	menu := &models.Menus{
		Name:    &params.Name,
		GroupId: &params.GroupId,
		Sort:    &params.Sort,
		Icon:    &params.Icon,
		Path:    &params.Path,
	}
	if res := global.DBMaster.Create(menu); res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
	}
	libs_http.RspState(ctx, 0, "创建成功")
}

func (c *Menu) Update(ctx *gin.Context) {
	// 接受参数
	var params struct {
		Id      uint64
		GroupId uint64
		Name    string
		Sort    int
		Icon    string
		Path    string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	menus := &models.Menus{}

	menu := &models.Menus{
		Id:      &params.Id,
		GroupId: &params.GroupId,
		Name:    &params.Name,
		Sort:    &params.Sort,
		Icon:    &params.Icon,
		Path:    &params.Path,
	}
	if err := menus.UpdateById(menu); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	libs_http.RspState(ctx, 0, "更新成功")
}

func (c *Menu) Del(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	menus := &models.Menus{
		Id: &params.Id,
	}
	err := menus.DeleteById(menus)
	if err != nil {
		libs_http.RspState(ctx, 0, err)
		return
	}

	libs_http.RspState(ctx, 0, "删除成功")
}
