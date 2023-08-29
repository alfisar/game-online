package controller

import "github.com/gin-gonic/gin"

type BankAccountControllerContract interface {
	Create(ctx *gin.Context)
}
