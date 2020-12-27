package tenant_admin_api

import (
	"api/config"
	libs_http "api/libs/http"
	"api/models"
	models_tenant "api/models/tenant"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Auth struct{}

func (c *Auth) Login(ctx *gin.Context) {
	// 接受参数
	var params struct {
		TenantCode string
		Username   string
		Password   string
		Code       string
		CodeId     string
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

	// 查询租户ID
	tenant := &models.Tenants{}
	{
		code := params.Code
		ok, err := tenant.SelectByCode("*", code)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		if !ok {
			libs_http.RspState(ctx, 1, "租户编码不存在")
			return
		}
	}

	admin := &models_tenant.Admins{}

	// 查询账号
	{
		tenantId := *tenant.Id
		username := params.Username
		ok, err := admin.SelectByUsername("*", tenantId, username)
		if err != nil {
			libs_http.RspState(ctx, 1, err)
			return
		}
		if !ok {
			libs_http.RspState(ctx, 1, "账号不存在")
			return
		}
	}

	// 校验用户
	{
		// 是否启用
		if *admin.Enable != true {
			libs_http.RspState(ctx, 1, "账号已经被禁用")
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
