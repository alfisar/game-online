package router

import "github.com/gin-gonic/gin"

func getRouterBank(router *gin.RouterGroup) {
	controller := setBankController()
	router.POST("/bank-account", setMiddleware().Authenticate(), controller.Create)

}
