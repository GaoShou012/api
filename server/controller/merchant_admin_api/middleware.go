package merchant_admin_api

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
	callback := &middleware_gin_redis_v8.Callback{
		Expiration: middleware_gin_redis_v8.Expiration(func(ctx *gin.Context) {
			libs_http.RspAuthFailed(ctx, 1, "登陆已经过期")
		}),
	}
	OperatorContext = middleware_gin_redis_v8.New(
		middleware_gin_redis_v8.WithModel(&Operator{}),
		middleware_gin_redis_v8.WithExpiration(time.Duration(conf.OperatorContextExpiration)*time.Second),
		middleware_gin_redis_v8.WithRedisClient(global.RedisClient),
		middleware_gin_redis_v8.WithCallback(callback),
		middleware_gin_redis_v8.WithCipherKey([]byte(conf.OperatorContextCipherKey)),
	)
}
