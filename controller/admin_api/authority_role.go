package admin_api

import (
	libs_http "api/libs/http"
	"api/models"
	"api/utils"
	"github.com/gin-gonic/gin"
)

type AuthorityRole struct{}

func (c *AuthorityRole) Get(ctx *gin.Context) {
	var params struct {
		Page int
		Size int
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &models.AuthorityRoles{}

	var list []models.AuthorityRoles
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

func (c *AuthorityRole) Create(ctx *gin.Context) {
	// 接受参数
	var params struct {
		Name   string
		Sort   int
		Remark string
		Enable bool
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &models.AuthorityRoles{
		Name:   &params.Name,
		Sort:   &params.Sort,
		Remark: &params.Remark,
		Enable: &params.Enable,
	}
	if res := utils.IMysql.Master.Create(model); res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
	}
	libs_http.RspState(ctx, 0, "创建成功")
}

func (c *AuthorityRole) Update(ctx *gin.Context) {
	// 接受参数
	var params struct {
		Id     uint64
		Name   string
		Sort   int
		Remark string
		Enable bool
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	model := &models.AuthorityRoles{
		Id:     &params.Id,
		Name:   &params.Name,
		Sort:   &params.Sort,
		Remark: &params.Remark,
		Enable: &params.Enable,
	}
	if err := model.UpdateById(model); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	libs_http.RspState(ctx, 0, "更新成功")
}

func (c *AuthorityRole) Del(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &models.AuthorityRoles{
		Id: &params.Id,
	}
	err := model.DeleteById(model)
	if err != nil {
		libs_http.RspState(ctx, 0, err)
		return
	}

	libs_http.RspState(ctx, 0, "删除成功")
}
