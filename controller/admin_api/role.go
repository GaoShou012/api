package admin_api

import (
	"api/libs/connect"
	libs_http "api/libs/http"
	"api/models"
	"github.com/gin-gonic/gin"
)

type Role struct{}

//创建角色
func (c *Role) CreateRole(ctx *gin.Context) {
	var params struct {
		Name     string
		Sequence int
		Memo     string
		Status   int
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	db := connect.GetDB()
	db.Create(&models.Role{
		Name:     &params.Name,
		Sequence: &params.Sequence,
		Memo:     &params.Memo,
		Status:   &params.Status,
	})
	libs_http.RspState(ctx,0,"创建成功")
}

//删除角色
func (c *Role) DeleteRole(ctx *gin.Context) {

}

//修改角色
func (c *Role) UpdateRole(ctx *gin.Context) {

}

//角色列表
func (c *Role) ListRole(ctx *gin.Context) {

}
