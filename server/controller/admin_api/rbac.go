package admin_api

import (
	"api/global"
	"github.com/gin-gonic/gin"
)

type Rbac struct {

}
func (c *Rbac) Menu(ctx *gin.Context){
	operator := GetOperator(ctx)


	// 查询所有角色
	roles := make([]string,0)
	{
		global.RBAC.SelectRoleByOperator(operator)
	}
}
