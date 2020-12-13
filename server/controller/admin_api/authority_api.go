package admin_api

import (
	libs_http "api/libs/http"
	"api/models"
	"github.com/gin-gonic/gin"
)

type AuthorityApi struct{}

func (c *AuthorityApi) Get(ctx *gin.Context) {
	var params struct {
		Page int
		Size int
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &models.AuthorityApis{}

	var list []models.AuthorityApis
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

func (c *AuthorityApi) Create(ctx *gin.Context) {
	// 接受参数
	var params struct {
		Method    string
		Path      string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &models.AuthorityApis{
		Method:   &params.Method,
		Path:   &params.Path,
	}
	if res := utils.IMysql.Master.Create(model); res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
	}
	libs_http.RspState(ctx, 0, "创建成功")
}

func (c *AuthorityApi) Update(ctx *gin.Context) {
	// 接受参数
	var params struct {
		Id     uint64
		Method    string
		Path      string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	model := &models.AuthorityApis{
		Id:     &params.Id,
		Method:   &params.Method,
		Path:   &params.Path,
	}
	if err := model.UpdateById(model); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	libs_http.RspState(ctx, 0, "更新成功")
}

func (c *AuthorityApi) Del(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &models.AuthorityApis{
		Id: &params.Id,
	}
	err := model.DeleteById(model)
	if err != nil {
		libs_http.RspState(ctx, 0, err)
		return
	}

	libs_http.RspState(ctx, 0, "删除成功")
}
