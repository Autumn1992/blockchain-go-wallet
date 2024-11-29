package api

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
}

var wr = WebRouter{}

func (r *Router) SetWalletRouter(engine *gin.Engine) {
	//暂时取消
	//engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.InstanceName(wallet.SwaggerInfoswagger_admin.InstanceName())))
	engine.POST("adminapi/doproxy", wr.DoProxy)
}

func NewRouter() *Router {
	router := &Router{}
	return router
}
