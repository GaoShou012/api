package merchant_admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"api/models/visitors"
	"github.com/gin-gonic/gin"
)

type VisitorManagement struct {
}

func (c *VisitorManagement) Update(ctx *gin.Context) {
	var params struct {
		//访客标签
		//Nickname       string
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
