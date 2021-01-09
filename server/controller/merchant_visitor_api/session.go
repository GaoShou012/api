package merchant_visitor_api

import (
	"api/cs"
	"api/cs/event"
	"api/cs/meta"
	"api/cs/notification"
	"api/global"
	libs_http "api/libs/http"
	libs_ip_location "api/libs/ip_location"
	"api/models"
	cs_env "cs/env"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Session struct{}

/*
	无需认证的接口
	访客进行会话之前，先创建会话

	POST

	@params
	Token 三方平台与客服平台，对称加密的用户信息
	Device 访客端的设备类型
*/
func (c *Session) Create(ctx *gin.Context) {
	var params struct {
		// 商户编码，用于识别customer token，需要使用指定的租户密钥，解密customer token
		MerchantCode string
		// 访客设备，手机，PC，等等
		VisitorDevice uint64
		// 访客Token，由租户生成，加密了访客的基本信息
		VisitorToken string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	var session *cs.Session
	merchant := &models.Merchants{}
	merchantSettings := &models.MerchantsSettings{}

	// 商户鉴权
	{
		exists, err := merchant.SelectByCode("*", params.MerchantCode)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		if !exists {
			libs_http.RspState(ctx, 1, fmt.Errorf("商户不存在"))
			return
		}
		if merchant.IsEnable() != true {
			libs_http.RspState(ctx, 1, fmt.Errorf("商户未启用"))
			return
		}
		if merchant.IsExpiration() {
			libs_http.RspState(ctx, 1, fmt.Errorf("商户租约已经过期"))
			return
		}
	}

	// 商户的客户Token密文本
	{
		exists, err := merchantSettings.SelectById("*", *merchant.Id)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		if !exists {
			libs_http.RspState(ctx, 1, fmt.Errorf("商户配置不存在"))
			return
		}
		if *merchantSettings.VisitorTokenCipherKey == "" {
			libs_http.RspState(ctx, 1, fmt.Errorf("商户的访客信息密本未设置"))
			return
		}
	}

	// 客服系统-客户端
	client := &meta.Client{
		MerchantId:   *merchant.Id,
		MerchantCode: *merchant.Code,
		UserId:       0,
		UserType:     uint64(meta.ClientTypeCustomer),
		Username:     "",
		Nickname:     "",
		Thumb:        "",
	}

	// 商户UserToken解析
	{
		visitor := &meta.Visitor{}
		token := params.VisitorToken
		cipherKey := *merchantSettings.VisitorTokenCipherKey
		if err := cs.CipherOfToken.DecryptWithCipherKey(token, visitor, []byte(cipherKey)); err != nil {
			libs_http.RspState(ctx, 1, fmt.Errorf("解析访客Token失败"))
			global.Logger.Error(err)
			return
		}
		client.UserId = visitor.UserId
		client.Username = visitor.Username
		client.Nickname = visitor.Nickname
		client.Thumb = visitor.Thumb
	}

	// 创建会话
	{
		customerIp := libs_http.GetIp(ctx)
		tmp, err := libs_ip_location.GetLocation(customerIp)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		customerIpLocation := fmt.Sprintf("%s:%s:%s", tmp.Country, tmp.Province, tmp.City)
		customerDevice := params.VisitorDevice
		s, err := cs.CustomerCreateSession(client, customerDevice, customerIp, customerIpLocation)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		session = s
	}

	// 返回数据
	rspData := make(map[string]interface{})

	// 会话信息
	{
		info := make(map[string]interface{})
		info["CreatedAt"] = session.CreatedAt

	}
	// 会话Token
	// 访客端操作API的OperatorContext
	{
		operator := &Operator{
			SessionId: session.Id,
			Client:    client,
			LoginTime: time.Now(),
		}
		token, err := OperatorContext.SignedString(operator)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		rspData["SessionToken"] = token
	}
	// 网关Token
	{
		token, err := cs.CipherOfToken.Encrypt(client)
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

	// 会话鉴权
	{
		sessionId := operator.SessionId
		session, err := cs.GetSessionInfo(sessionId)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		if session.GetEnable() == false {
			libs_http.RspState(ctx, 1, "会话已经关闭")
			return
		}
	}

	// 会话消息
	{
		sessionId := operator.SessionId
		client := operator.Client
		content := params.Content
		contentType := params.ContentType
		messageId, err := cs.SessionMessage(sessionId, client, content, contentType)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		libs_http.RspData(ctx, 0, "", messageId)
		return
	}
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

	// 会话ID，评分，评语
	sessionId := operator.SessionId
	rating := params.Rating
	comment := params.Comment

	// 标记会话已经评价
	{
		model := &models.MerchantsSessions{}
		if err := model.Rating(sessionId, rating, comment); err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
	}

	// 推送系统消息到会话消息流
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
