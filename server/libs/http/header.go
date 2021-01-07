package libs_http

import (
	"github.com/gin-gonic/gin"
	"strings"
)

/*
	Gin 框架中获取http请求的客服端IP
*/
func GetIp(ctx *gin.Context) string {
	addr := ctx.Request.RemoteAddr
	arr := strings.Split(addr, ":")
	return arr[0]
}
