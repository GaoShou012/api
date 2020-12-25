package admin_api

import (
	"api/global"
	libs_http "api/libs/http"
	"framework/class/middleware"
	"framework/plugin/middleware/middleware_gin"
	"github.com/gin-gonic/gin"
)

var OperatorContext middleware.OperatorContext

func GetOperator(ctx *gin.Context) (*Operator, error) {
	operator, err := OperatorContext.Get(ctx)
	if err != nil {
		return nil, err
	}
	return operator.(*Operator), err
}

func InitOperatorContext() {
	callback := &middleware_gin.Callback{
		Expiration: func(ctx *gin.Context) {
			libs_http.RspAuthFailed(ctx, 1, "登陆已经过期")
		},
	}
	OperatorContext = middleware_gin.New(
		middleware_gin.WithModel(&Operator{}),
		middleware_gin.WithRedisClient(global.RedisClient),
		middleware_gin.WithCallback(callback),
		middleware_gin.WithCipherKey([]byte("123123")),
	)
}
