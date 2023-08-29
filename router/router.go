package router

import (
	"game-online/config"

	"github.com/gin-gonic/gin"
)

var (
	_   = config.SetConfig()
	cfg = config.GetConfig()
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	getRouterUser(api)
	getRouterWallet(api)
	getRouterBank(api)
	return router
}
