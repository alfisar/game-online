package controller

import "github.com/gin-gonic/gin"

type WalletControllerContract interface {
	TopUp(ctx *gin.Context)
}
