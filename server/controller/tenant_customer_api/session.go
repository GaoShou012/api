package tenant_customer_api

import (
	"api/cs"
	"api/cs/message"
	"api/cs/notification"
	"api/global"
	libs_http "api/libs/http"
	"api/meta"
	"api/models"
	cs_env "cs/env"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type Session struct {
}

// 客户创建会话
func (c *Session) Create(ctx *gin.Context) {

	// 查询租户是否开通

	client := &meta.Client{
		TenantCode: "",
		UserId:     0,
		Username:   "",
		UserType:   "",
	}
	session := &meta.Session{Id: uuid.NewV1().String()}

	// 创建会话
	err := global.CsSys.CreateSession(client, session)
	if err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	// 会话排队

	// 广播会话数量给所有在线客服
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
		msg := cs.NewMessageWithContent(operator,&notification.SessionRating{
			SessionId: "",
			Rating:    0,
			Comment:   "",
		})
		encode, err := message.Encode(msg)
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
