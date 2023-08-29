package router

import "github.com/gin-gonic/gin"

func getRouterWallet(router *gin.RouterGroup) {
	controller := setWalletController()
	router.POST("/wallet/topup", setMiddleware().Authenticate(), controller.TopUp)

}
