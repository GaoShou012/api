package tenant_customer_api

import (
	"api/config"
	"api/global"
	libs_http "api/libs/http"
	"fmt"
	"framework/class/middleware"
	"framework/plugin/middleware/middleware_gin"
	"github.com/gin-gonic/gin"
	"time"
)

var OperatorContext middleware.OperatorContext

func GetOperator(ctx *gin.Context) *Operator {
	operator, err := OperatorContext.Get(ctx)
	if err != nil {
		panic(fmt.Errorf("GetOperator is failed err=%s", err))
	}
	return operator.(*Operator)
}

func InitOperatorContext() {
	conf := config.GetConfig().Base
	callback := &middleware_gin.Callback{
		Expiration: middleware_gin.Expiration(func(ctx *gin.Context) {
			libs_http.RspAuthFailed(ctx, 1, "登陆已经过期")
		}),
	}
	OperatorContext = middleware_gin.New(
		middleware_gin.WithModel(&Operator{}),
		middleware_gin.WithExpiration(time.Duration(conf.OperatorContextExpiration)*time.Second),
		middleware_gin.WithRedisClient(global.RedisClient),
		middleware_gin.WithCallback(callback),
		middleware_gin.WithCipherKey([]byte(conf.OperatorContextCipherKey)),
	)
}
