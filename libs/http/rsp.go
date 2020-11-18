package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	Description: 	用于http请求的统一返回
	Developer: 		高手
*/

const (
	RspTypeCommon = 0
	RspTypeAuth   = 1
)

/*
	返回验证失败专用
*/
func RspAuthFailed(ctx *gin.Context, code int, desc interface{}) {
	res := map[string]interface{}{
		"Type": RspTypeAuth,
		"Code": code,
		"Desc": AssertDesc(desc),
		"Data": "",
	}
	ctx.JSON(http.StatusOK, res)
}

/*
	通用返回
*/
func Rsp(ctx *gin.Context, code int, desc interface{}, data interface{}) {
	res := map[string]interface{}{
		"Type": RspTypeCommon,
		"Code": code,
		"Desc": AssertDesc(desc),
		"Data": data,
	}
	ctx.JSON(http.StatusOK, res)
}

/*
	通用返回状态
	默认Data是空字符串
*/
func RspState(ctx *gin.Context, code int, desc interface{}) {
	res := map[string]interface{}{
		"Type": RspTypeCommon,
		"Code": code,
		"Desc": AssertDesc(desc),
		"Data": "",
	}
	ctx.JSON(http.StatusOK, res)
}

/*
	通用返回数据
	Data泛型类型，由业务层保持非nil
*/
func RspData(ctx *gin.Context, code int, desc interface{}, data interface{}) {
	res := map[string]interface{}{
		"Type": RspTypeCommon,
		"Code": code,
		"Desc": AssertDesc(desc),
		"Data": data,
	}
	ctx.JSON(http.StatusOK, res)
}

/*
	通用返回查询数据
	rows泛型类型，可以由此方法保持空数组
*/
func RspSearch(ctx *gin.Context, code int, desc interface{}, total int64, rows interface{}) {
	if rows == nil {
		rows = make([]interface{}, 0)
	}
	data := struct {
		Total int64
		List  interface{}
	}{Total: total, List: rows}
	res := map[string]interface{}{
		"Type": RspTypeCommon,
		"Code": code,
		"Desc": AssertDesc(desc),
		"Data": data,
	}
	ctx.JSON(http.StatusOK, res)
}

func AssertDesc(desc interface{}) string {
	var descString string
	switch desc.(type) {
	case string:
		descString = desc.(string)
	case error:
		descString = desc.(error).Error()
	case nil:
		descString = ""
		break
	default:
		descString = "Assert desc failed"
	}
	return descString
}
