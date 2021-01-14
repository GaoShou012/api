package merchant_visitor_api

import (
	"github.com/gin-gonic/gin"
)

type HttpService struct{}

func (r *HttpService) Cors() gin.HandlerFunc {
	return nil
}

func (r *HttpService) Route(engine *gin.Engine) {
}
