package merchant_customer_api

import "github.com/gin-gonic/gin"

type Merchant struct {
}

/*
	访客访问客户系统时
	商户的欢迎语
*/
func (c *Merchant) Welcome(ctx *gin.Context) {
	operator := GetOperator(ctx)

}
