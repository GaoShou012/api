package admin_api

import (
	libs_http "api/libs/http"
	"api/models"
	"github.com/gin-gonic/gin"
)

type Casbin struct{}

//添加权限
func (c *Casbin) AddPolicy(ctx *gin.Context) {
	var params struct {
		Path        string
		AuthorityId string
		Method      string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	m := models.CasbinRule{}
	rules := [][]string{}
	rules = append(rules, []string{params.AuthorityId, params.Path, params.Method})
	err := m.AddCasbinPolicy(rules)
	if err != nil {
		libs_http.RspState(ctx, 1, err)
	}
	libs_http.RspState(ctx, 0, "创建成功")
	return
}

//删除权限
func (c *Casbin) DelPolicy(ctx *gin.Context) {

}

//修改权限
func (c *Casbin) UpdatePolicy(ctx *gin.Context) {

}
