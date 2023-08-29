package router

import (
	"github.com/gin-gonic/gin"
)

func getRouterUser(router *gin.RouterGroup) {
	controller := setUserController()
	router.POST("/user/login", controller.Login)
	router.POST("/user/logout", setMiddleware().Authenticate(), controller.Logout)
	router.POST("/user/register", controller.Register)
	router.GET("/user/list", setMiddleware().Authenticate(), controller.GetList)
	router.GET("/user/detail", setMiddleware().Authenticate(), controller.Detail)
}
