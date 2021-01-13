package merchant_admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"api/models"
	"github.com/gin-gonic/gin"
)

type IpWhitelist struct {
}

func (c *IpWhitelist) Get(ctx *gin.Context) {
	var params struct {
		Page       int
		PageSize   int
		MerchantId uint64
		Ip         string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &models.IpWhitelist{}
	data := &[]models.IpWhitelist{}
	//var where string
	var total int64
	db := global.DBSlave.Table(model.GetTableName())
	if params.MerchantId < 0 {
		//where = fmt.Sprintf("merchant_id = %d",params.MerchantId)
		db.Where("merchant_id = ?", params.MerchantId)
	}

	if params.Ip != "" {
		db.Where("ip = ?", params.Ip)
	}
	db.Limit(params.Page).Offset((params.PageSize - 1) * params.PageSize)
	db.Order("create_at desc")
	db.Find(data).Count(&total)
	libs_http.RspSearch(ctx, 0, "请求成功", total, data)
	return
}
func (c *IpWhitelist) Create(ctx *gin.Context) {
	var params struct {
		MerchantId uint64 //商户id
		Ip         string //白名单ip
		Desc       string //描述
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	whitelist := &models.IpWhitelist{
		MerchantId: &params.MerchantId,
		Ip:         &params.Ip,
		Desc:       &params.Desc,
	}
	res := global.DBMaster.Table(whitelist.GetTableName()).Create(whitelist)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "添加成功")
	return
}
func (c *IpWhitelist) Update(ctx *gin.Context) {
	var params struct {
		Id         uint64
		MerchantId uint64 //商户id
		Ip         string //白名单ip
		Desc       string //描述
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	whitelist := &models.IpWhitelist{
		Model:      models.Model{Id: &params.Id},
		MerchantId: &params.MerchantId,
		Ip:         &params.Ip,
		Desc:       &params.Desc,
	}
	res := global.DBMaster.Table(whitelist.GetTableName()).Update(whitelist)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "更新成功")
	return
}
func (c *IpWhitelist) Delete(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	whitelist := &models.IpWhitelist{
		Model: models.Model{Id: &params.Id},
	}
	res := global.DBMaster.Table(whitelist.GetTableName()).Delete(whitelist)
	if res.Error != nil {
		libs_http.RspState(ctx, 1, res.Error)
		return
	}
	libs_http.RspState(ctx, 0, "删除成功")
	return
}
