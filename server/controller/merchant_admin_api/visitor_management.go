package merchant_admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"api/models/visitors"
	"github.com/gin-gonic/gin"
)

type VisitorManagement struct {
}

func (c *VisitorManagement) VisitorList(ctx *gin.Context) {
	var params struct {
		Page       uint
		PageSize   uint
		MerchantId uint64
		Username   string
		StartAt    string
		EndAt      string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	model := &visitors.Visitors{}
	data := &[]visitors.Visitors{}
	//var where string
	var total int64
	db := global.DBSlave.Table(model.GetTableName())
	if params.MerchantId < 0 {
		//where = fmt.Sprintf("merchant_id = %d",params.MerchantId)
		db.Where("merchant_id = ?", params.MerchantId)
	}
	if params.Username != "" {
		//where = fmt.Sprintf(where+"and"+"username = %s",params.Username)
		db.Where("username = ?", params.MerchantId)
	}
	if params.StartAt != "" && params.EndAt != "" {
		db.Where("create_at between ? and ?", params.StartAt, params.EndAt)
	}
	db.Limit(params.Page).Offset((params.PageSize - 1) * params.PageSize)
	db.Order("create_at")
	db.Find(data).Count(&total)
	//total, err := model.Count(where)
	//if err != nil {
	//	libs_http.RspState(ctx, 1, err)
	//	return
	//}
	libs_http.RspSearch(ctx, 0, "请求成功", total, data)
	return
}

func (c *VisitorManagement) Update(ctx *gin.Context) {
	var params struct {
		Tags           string
		Gender         int
		Phone          uint64
		Email          string
		Wechat         string
		WechatNickname string
		QQ             string
		QQNickname     string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	visitor := &visitors.Visitors{
		Tags:           &params.Tags,
		Gender:         &params.Gender,
		Phone:          &params.Phone,
		Email:          &params.Email,
		Wechat:         &params.Wechat,
		WechatNickname: &params.WechatNickname,
		QQ:             &params.QQ,
		QQNickname:     &params.QQNickname,
	}

	if err := global.DBMaster.Update(visitor); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	libs_http.RspState(ctx, 0, "修改成功")
	return
}
