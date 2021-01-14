package merchant_api

import (
	controller "api/controller/merchant_admin_api"
	"github.com/gin-gonic/gin"
)

type HttpService struct{}

func (r *HttpService) Cors() gin.HandlerFunc {
	return nil
}

func (r *HttpService) Route(engine *gin.Engine) {
	controller.InitOperatorContext()

	//api := engine.Group(fmt.Sprintf("/merchant%s", meta.ApiVersion))
	//var authenticated gin.IRoutes
}
