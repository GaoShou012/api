package libs_http

import (
	"api/global"
	"github.com/gin-gonic/gin"
)

func ParsePostJson(ctx *gin.Context, v interface{}) bool {
	if err := ctx.BindJSON(v); err != nil {
		RspState(ctx, 1, "解析参数失败")
		global.Logger.Error(err)
		return false
	} else {
		return true
	}
}
func ParseGet(ctx *gin.Context, v interface{}) bool {
	if err := ctx.Bind(v); err != nil {
		RspState(ctx, 1, "解析参数失败")
		global.Logger.Error(err)
		return false
	} else {
		return true
	}
}
