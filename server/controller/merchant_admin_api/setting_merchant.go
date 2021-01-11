package merchant_admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"api/models"
	"github.com/gin-gonic/gin"
)

type SettingMerchant struct {
}

func (c *SettingMerchant) Get(ctx *gin.Context) {

}
func (c *SettingMerchant) Create(ctx *gin.Context) {
	var params struct {
		Type                 int    //类型 1pc 2手机端
		Logo                 string //商户logo
		CustomerServiceImage string //客服头像
		VisitorImage         string //访客头像
		LeftAd               string //左侧广告
		LeftAdUrl            string //左侧广告连接
		RightAd              string //右侧广告
		RightAdUrl           string //右侧广告连接
		Color                string //窗口颜色配置
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	settingMerchant := models.SettingMerchants{
		//Model:                models.Model{},
		Type:                 &params.Type,
		Logo:                 &params.Logo,
		CustomerServiceImage: &params.CustomerServiceImage,
		VisitorImage:         &params.VisitorImage,
		LeftAd:               &params.LeftAd,
		LeftAdUrl:            &params.LeftAdUrl,
		RightAd:              &params.RightAd,
		RightAdUrl:           &params.LeftAdUrl,
		Color:                &params.Color,
	}
	res := global.DBMaster.Table(settingMerchant.GetTableName()).Create(settingMerchant)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "添加成功")
}
func (c *SettingMerchant) Update(ctx *gin.Context) {
	var params struct {
		Id                   uint64
		Type                 int    //类型 1pc 2手机端
		Logo                 string //商户logo
		CustomerServiceImage string //客服头像
		VisitorImage         string //访客头像
		LeftAd               string //左侧广告
		LeftAdUrl            string //左侧广告连接
		RightAd              string //右侧广告
		RightAdUrl           string //右侧广告连接
		Color                string //窗口颜色配置
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	var settingMerchant models.SettingMerchants
	settingMerchant.Id = &params.Id
	settingMerchant.Type = &params.Type
	settingMerchant.Logo = &params.Logo
	settingMerchant.CustomerServiceImage = &params.CustomerServiceImage
	settingMerchant.VisitorImage = &params.VisitorImage
	settingMerchant.LeftAd = &params.LeftAd
	settingMerchant.LeftAdUrl = &params.LeftAdUrl
	settingMerchant.RightAd = &params.RightAd
	settingMerchant.RightAdUrl = &params.RightAdUrl
	settingMerchant.Color = &params.Color
	res := global.DBMaster.Table(settingMerchant.GetTableName()).Update(&settingMerchant)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "修改成功")
	return
}
func (c *SettingMerchant) Delete(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	var settingMerchant models.SettingMerchants
	settingMerchant.Id = &params.Id
	res := global.DBMaster.Table(settingMerchant.GetTableName()).Update(&settingMerchant)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "删除成功")
	return
}
