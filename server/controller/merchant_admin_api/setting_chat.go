package merchant_admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"api/models"
	"github.com/gin-gonic/gin"
)

type SettingChat struct {
}

func (c *SettingChat) Get(ctx *gin.Context) {

}
func (c *SettingChat) Create(ctx *gin.Context) {
	var params struct {
		AssignRule          int
		IsAssignLastService bool
		RobotEnable         bool
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	settingChat := &models.SettingChats{
		AssignRule:          &params.AssignRule,
		IsAssignLastService: &params.IsAssignLastService,
		RobotEnable:         &params.RobotEnable,
	}
	res := global.DBMaster.Table(settingChat.GetTableName()).Create(settingChat)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "添加成功")
	return
}
func (c *SettingChat) Update(ctx *gin.Context) {
	var params struct {
		Id                  uint64
		AssignRule          int
		IsAssignLastService bool
		RobotEnable         bool
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	var settingChat models.SettingChats
	settingChat.Id = &params.Id
	settingChat.IsAssignLastService = &params.IsAssignLastService
	settingChat.AssignRule = &params.AssignRule
	settingChat.RobotEnable = &params.RobotEnable

	res := global.DBMaster.Table(settingChat.GetTableName()).Update(&settingChat)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "修改成功")
	return
}
func (c *SettingChat) Delete(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	var settingChat models.SettingChats
	settingChat.Id = &params.Id
	res := global.DBMaster.Table(settingChat.GetTableName()).Delete(&settingChat)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "删除成功")
	return
}
