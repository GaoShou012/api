package admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"api/models"
	"github.com/gin-gonic/gin"
)

type MenuGroup struct{}

func (c *MenuGroup) Get(ctx *gin.Context) {
	var params struct {
		Page int
		Size int
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &models.MenusGroups{}

	var list []models.Menus
	// 按名称查找
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
	libs_http.RspSearch(ctx, 0, nil, count, list)
	return
}

func (c *MenuGroup) Create(ctx *gin.Context) {
	// 接受参数
	var params struct {
		Sort  uint64
		Group string
		Icon  string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	menuGroup := &models.MenusGroups{
		Sort:      &params.Sort,
		GroupName: &params.Group,
		Icon:      &params.Icon,
	}
	if res := global.DBMaster.Create(menuGroup); res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
	}
	libs_http.RspState(ctx, 0, "创建成功")
}

func (c *MenuGroup) Update(ctx *gin.Context) {
	// 接受参数
	var params struct {
		Id        uint64
		Sort      uint64
		GroupName string
		Icon      string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	//model := &models.MenusGroups{}

	model := &models.MenusGroups{
		Id:        &params.Id,
		GroupName: &params.GroupName,
		Sort:      &params.Sort,
		Icon:      &params.Icon,
	}
	if err := model.UpdateById(model); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	libs_http.RspState(ctx, 0, "更新成功")
}

func (c *MenuGroup) Del(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &models.MenusGroups{
		Id: &params.Id,
	}
	err := model.DeleteById(model)
	if err != nil {
		libs_http.RspState(ctx, 0, err)
		return
	}

	libs_http.RspState(ctx, 0, "删除成功")
}
