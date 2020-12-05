package admin_api

import (
	"api/config"
	libs_http "api/libs/http"
	"api/meta"
	"api/models"
	"api/utils"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"golang.org/x/crypto/bcrypt"
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

func (c *Auth) Register(ctx *gin.Context) {
	// 接受参数
	var params struct {
		Username string
		Password string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	{
		var admin models.Admins
		db := utils.IMysql.Slave
		db.Where("username = ?", params.Username).First(&admin)
		if admin.ID != nil {
			libs_http.RspState(ctx, 1, "用户已经存在")
			return
		}

		//密码加密
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
		if err != nil {
			libs_http.RspState(ctx, 1, "加密错误")
			return
		}
		params.Password = string(hashPassword)
		db.Create(&models.Admins{
			Username: &params.Username,
			Password: &params.Password,
		})
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
		var admin models.Admins
		db := utils.IMysql.Slave
		db.Where("username = ?", params.Username).First(&admin)
		if admin.ID == nil {
			libs_http.RspState(ctx, 1, "用户不存在")
			return
		}
		//验证密码
		if err := bcrypt.CompareHashAndPassword([]byte(*admin.Password), []byte(params.Password)); err != nil {
			libs_http.RspState(ctx, 1, "密码错误")
			return
		}
		//db.Create(&models.Admins{
		//go get -u golang.org/x/crypto/bcrypt
		//	Username: &params.Username,
		//	Password: &params.Password,
		//	Nickname: &params.Username,
		//})

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
	libs_http.RspData(ctx, 123, nil, "exit success")
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
