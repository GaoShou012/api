package admin_api

import (
	"api/config"
	libs_http "api/libs/http"
	"api/meta"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type Auth struct{}

func (c *Auth) Verify(ctx *gin.Context) {
	// 获取Token
	token := ctx.GetHeader(meta.XApiToken)

	// 获取解析Token密钥
	key := config.GetConfig().Base.TokenKey

	// 解析Token
	operator := &Operator{}
	if err := operator.decrypt([]byte(key), token); err != nil {
		libs_http.RspAuthFailed(ctx, 1, err)
		ctx.Abort()
		return
	}

	// 保存Token到上下文
	SetOperator(ctx, operator)
}

func (c *Auth) Login(ctx *gin.Context) {
	// 接受参数
	var params struct {
		Username string
		Password string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	// 校验用户
	{
		
	}

	// 生成Token
	{
		conf := config.GetConfig().Base

		// TODO 赋值相应的数据
		operator := Operator{}

		token, err := operator.encrypt([]byte(conf.TokenKey))
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		libs_http.RspData(ctx, 0, nil, token)
	}
}

func (c *Auth) Logout(ctx *gin.Context) {

}

func (c *Auth) CodeImage(ctx *gin.Context) {
	captcha := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, base64Captcha.DefaultMemStore)
	id, b64s, err := captcha.Generate()
	if err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	res := make(map[string]interface{})
	res["id"] = id
	res["image"] = b64s
	libs_http.RspData(ctx, 0, nil, res)
}

/*
	校验验证码是否正确
	@return
	true 正确
	false 不正确
*/
func (c *Auth) verifyCode(id string, code string) bool {
	return base64Captcha.DefaultMemStore.Verify(id, code, true)
}
