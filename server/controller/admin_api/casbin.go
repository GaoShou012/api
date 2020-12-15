package admin_api

import (
	libs_http "api/libs/http"
	"api/models"
	"fmt"
	"github.com/casbin/casbin/v2/util"
	"github.com/gin-gonic/gin"
	"strings"
)

type Casbin struct {}
// 拦截器
func(c *Casbin) CasbinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		operator, _ := GetOperator(ctx)
		AuthorityId := operator.Username
		// 获取请求的URI
		obj := ctx.Request.URL.RequestURI()
		// 获取请求方法
		act := ctx.Request.Method
		// 获取用户的角色
		sub := AuthorityId
		fmt.Println(sub,obj,act)

		casbinModel := models.CasbinRule{}
		success,err := casbinModel.ExecutePermission(sub, obj, act)
		if success {
			ctx.Next()
		} else {
			libs_http.RspState(ctx, 1, err)
			ctx.Abort()
			return
		}
	}
}

func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return ParamsMatch(name1, name2), nil
}

func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}