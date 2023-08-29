package controller

import "github.com/gin-gonic/gin"

type UserControllerContract interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}
