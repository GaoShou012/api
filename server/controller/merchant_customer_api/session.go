package merchant_customer_api

import (
	"api/cs"
	"api/cs/gateway"
	"api/cs/notification"
	"api/global"
	libs_http "api/libs/http"
	libs_ip_location "api/libs/ip_location"
	"api/meta"
	"api/models"
	cs_env "cs/env"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type Session struct{}

// 客户创建会话
func (c *Session) Create(ctx *gin.Context) {
	var params struct {
		// 租户编码，用于识别customer token，需要使用指定的租户密钥，解密customer token
		TenantCode string
		// 访客设备，手机，PC，等等
		CustomerDevice uint64
		// 访客Token，由租户生成，加密了访客的基本信息
		CustomerToken string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	// TODO 查询租户是否开通
	{

	}

	client := &meta.Client{
		TenantCode: "",
		UserId:     0,
		Username:   "",
		UserType:   "",
	}
	session := &meta.Session{Id: uuid.NewV1().String()}

	// 创建会话
	{
		customerIp := libs_http.GetIp(ctx)
		tmp, err := libs_ip_location.GetLocation(customerIp)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		customerIpLocation := fmt.Sprintf("%s:%s:%s", tmp.Country, tmp.Province, tmp.City)
		customerDevice := params.CustomerDevice
		session, err := cs.CustomerCreateSession(client, customerDevice, customerIp, customerIpLocation)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
	}

	// 返回数据
	rspData := make(map[string]interface{})

	// 会话token
	{
		token, err := cs.CipherOfToken.Encrypt(session)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		rspData["SessionToken"] = token
	}

	// 网关token
	{
		gatewayToken := gateway.Token{
			TenantCode: "",
			UserType:   "",
			UserId:     0,
		}
		token, err := cs.CipherOfToken.Encrypt(gatewayToken)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		rspData["GatewayToken"] = token
	}

	libs_http.RspData(ctx, 0, "", rspData)
}

/*
	客户发送会话消息
*/
func (c *Session) Message(ctx *gin.Context) {
	var params struct {
		Content     string
		ContentType string
	}

	operator := GetOperator(ctx)

	// 读取会话信息
	// 会话鉴权
	session := &cs.Session{}
	if err := cs_env.Session.GetInfo(operator.SessionId, session); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	if session.IsClose() {
		libs_http.RspState(ctx, 1, "会话已经关闭")
		return
	}

	// 广播消息
	messageId, err := global.CsSys.Broadcast(operator, operator.SessionId, params)
	if err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	libs_http.RspData(ctx, 0, "", messageId)
}

/*
	客户对会话服务进行评价
*/
func (c *Session) Rating(ctx *gin.Context) {
	var params struct {
		Rating  uint64
		Comment string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	// 获取操作者信息
	operator := GetOperator(ctx)

	sessionId := operator.SessionId
	rating := params.Rating
	comment := params.Comment

	// 标记会话已经评价
	{
		model := &models.TenantsSessions{}
		if err := model.Rating(sessionId, rating, comment); err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
	}

	// 通知客服，会话已经被评价
	go func() {
		msg := cs.NewMessageWithContent(operator, &notification.SessionRating{
			SessionId: "",
			Rating:    0,
			Comment:   "",
		})
		encode, err := event.Encode(msg)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}

		if err := cs_env.Gateway.Publish(operator.GetUUID(), encode); err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
	}()

	libs_http.RspState(ctx, 0, "评价成功")
}
