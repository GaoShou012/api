package admin_api

import (
	"api/config"
	"api/global"
	libs_http "api/libs/http"
	"api/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Auth struct{}

func (c *Auth) Register(ctx *gin.Context) {
	// 接受参数
	var params struct {
		Enable   bool
		State    uint64
		UserType uint64
		Username string
		Nickname string
		Password string

		Code   string
		CodeId string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	// 校验验证码
	{
		id, code := params.CodeId, params.Code
		if c.verifyCode(id, code) == false {
			libs_http.RspState(ctx, 1, errors.New("验证码错误"))
			return
		}
	}

	// 查询用户是否存在
	{
		admin := &models.Admins{}
		exists, err := admin.IsExistsByUsername(params.Username)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		if exists {
			libs_http.RspState(ctx, 1, "用户已经存在")
			return
		}
	}

	var admin *models.Admins

	{
		//密码加密
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
		if err != nil {
			libs_http.RspState(ctx, 1, "加密错误")
			return
		}
		encryptPassword := string(hashPassword)

		admin = &models.Admins{
			//Id:        nil,
			Enable:   &params.Enable,
			State:    &params.State,
			UserType: &params.UserType,
			Username: &params.Username,
			Password: &encryptPassword,
			Nickname: &params.Nickname,
			//CreatedAt: nil,
			//UpdatedAt: nil,
		}

		if res := global.DBMaster.Create(admin); res.Error != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
	}

	// 生成Token
	{
		// TODO 赋值相应的数据
		operator := &Operator{
			UserId:    *admin.Id,
			UserType:  *admin.UserType,
			Username:  *admin.Username,
			Nickname:  *admin.Nickname,
			LoginTime: time.Now(),
		}

		token, err := OperatorContext.SignedString(operator)
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
		Code     string
		CodeId   string
	}
	if err := ctx.BindJSON(&params); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}

	// 校验验证码
	{
		id, code := params.CodeId, params.Code
		if c.verifyCode(id, code) == false {
			libs_http.RspState(ctx, 1, errors.New("验证码错误"))
			return
		}
	}

	admin := &models.Admins{}

	// 校验用户
	{
		if err := admin.SelectByUsername("*", params.Username); err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		//验证密码
		if err := bcrypt.CompareHashAndPassword([]byte(*admin.Password), []byte(params.Password)); err != nil {
			libs_http.RspState(ctx, 1, "密码错误")
			return
		}
	}

	// 生成Token
	{
		// TODO 赋值相应的数据
		operator := &Operator{
			UserId:    *admin.Id,
			UserType:  *admin.UserType,
			Username:  *admin.Username,
			Nickname:  *admin.Nickname,
			LoginTime: time.Now(),
		}

		token, err := OperatorContext.SignedString(operator)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		libs_http.RspData(ctx, 0, nil, token)
	}
}

func (c *Auth) Logout(ctx *gin.Context) {
	if err := OperatorContext.Release(ctx); err != nil {
		libs_http.RspState(ctx, 1, err)
		return
	}
	libs_http.RspState(ctx, 0, "退出成功")
}

/*
	获取验证码
*/
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
	// 如果不开启验证码，不管输入任意内容，直接返回true
	if config.GetConfig().EnableAuthCode == false {
		return true
	}
	if id == "" || code == "" {
		return false
	}
	return base64Captcha.DefaultMemStore.Verify(id, code, true)
}
