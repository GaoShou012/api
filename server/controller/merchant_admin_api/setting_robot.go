package merchant_admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"api/models"
	"github.com/gin-gonic/gin"
)

type SettingRobot struct {
}

func (m *SettingRobot) Get(ctx *gin.Context) {

}
func (m *SettingRobot) Create(ctx *gin.Context) {
	var params struct {
		AutoClose    uint
		AfterAskTime uint
		WaitAsk      uint
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	settingRobots := &models.SettingRobots{
		AutoClose:    &params.AutoClose,
		AfterAskTime: &params.AfterAskTime,
		WaitAsk:      &params.WaitAsk,
	}
	res := global.DBMaster.Table(settingRobots.GetTableName()).Create(settingRobots)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "添加成功")
	return
}
func (m *SettingRobot) Update(ctx *gin.Context) {
	var params struct {
		Id           uint64
		AutoClose    uint
		AfterAskTime uint
		WaitAsk      uint
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	settingRobots := &models.SettingRobots{
		Model:        models.Model{Id: &params.Id},
		AutoClose:    &params.AutoClose,
		AfterAskTime: &params.AfterAskTime,
		WaitAsk:      &params.WaitAsk,
	}
	res := global.DBMaster.Table(settingRobots.GetTableName()).Update(settingRobots)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "修改成功")
	return
}
func (m *SettingRobot) Delete(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	settingRobots := &models.SettingRobots{
		Model: models.Model{Id: &params.Id},
	}
	res := global.DBMaster.Table(settingRobots.GetTableName()).Delete(settingRobots)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "删除成功")
	return
}
