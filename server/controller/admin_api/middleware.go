package admin_api

import (
	"api/config"
	"api/global"
	libs_http "api/libs/http"
	"fmt"
	"framework/class/middleware"
	"framework/plugin/middleware/middleware_gin_redis"
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
	callback := &middleware_gin_redis.Callback{
		Expiration: middleware_gin_redis.Expiration(func(ctx *gin.Context) {
			libs_http.RspAuthFailed(ctx, 1, "登陆已经过期")
		}),
		Error: func(ctx *gin.Context, err error) {
			libs_http.RspAuthFailed(ctx, 1, err)
		},
	}
	OperatorContext = middleware_gin_redis.New(
		middleware_gin_redis.WithModel(&Operator{}),
		middleware_gin_redis.WithExpiration(time.Duration(conf.OperatorContextExpiration)*time.Second),
		middleware_gin_redis.WithRedisClient(global.RedisClient),
		middleware_gin_redis.WithCallback(callback),
		middleware_gin_redis.WithCipherKey([]byte(conf.OperatorContextCipherKey)),
	)
}
