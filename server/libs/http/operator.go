package libs_http

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type Operator interface {
	GetTenantId() uint64
	GetAuthorityId() string
}

func SetOperator(ctx *gin.Context, operator Operator) {
	ctx.Set("operator", operator)
}
func GetOperator(ctx *gin.Context) (Operator, error) {
	val, exists := ctx.Get("operator")
	if !exists {
		return nil, errors.New("Lose the operator info\n")
	}
	op, ok := val.(Operator)
	if !ok {
		return nil, errors.New("Assert operator type failed\n")
	}
	return op, nil
}
