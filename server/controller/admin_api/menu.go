package admin_api

import (
	libs_http "api/libs/http"
	"api/models"
	"api/utils"
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
	data := make(map[string]interface{})
	// 按名称查找
	if params.Name != "" {
		filed := fmt.Sprintf("`name` LIKE '%%%s%%'", params.Name)
		count, err := menus.Count(filed)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		res := utils.IMysql.Slave.Order("id desc").Limit(params.Size)
		res.Offset(params.Page * params.Size).Where(filed).Find(&list)
		if res.Error != nil {
			libs_http.RspState(ctx, 0, res.Error)
			return
		}
		data["count"] = count
		data["data"] = list
		libs_http.RspData(ctx, 0, nil, data)
		return
	}
	// 按名称查找
	count, err := menus.Count("*")
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

func (c *Menu) Create(ctx *gin.Context) {
	// 接受参数
	var params struct {
		Name string
		Sort int
		Icon string
		Path string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	menu := &models.Menus{
		Name: &params.Name,
		Sort: &params.Sort,
		Icon: &params.Icon,
		Path: &params.Path,
	}
	if res := utils.IMysql.Master.Create(menu); res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
	}
	libs_http.RspState(ctx, 0, "创建成功")
}

func (c *Menu) Update(ctx *gin.Context) {
	// 接受参数
	var params struct {
		Id   uint64
		Name string
		Sort int
		Icon string
		Path string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	menus := &models.Menus{}

	menu := &models.Menus{
		Id:   &params.Id,
		Name: &params.Name,
		Sort: &params.Sort,
		Icon: &params.Icon,
		Path: &params.Path,
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