package middleware_gin_redis_v8

import "github.com/gin-gonic/gin"

type Callback struct {
	Expiration
}

/*
	当检测到上下文信息，已经过期时，进行回调
*/
type Expiration func(ctx *gin.Context)
