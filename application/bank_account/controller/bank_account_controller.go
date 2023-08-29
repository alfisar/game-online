package controller

import (
	"fmt"
	"game-online/application/bank_account/service"
	"game-online/domain"
	"game-online/internal/errorhandler"
	"game-online/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BankAccountController struct {
	serv service.BankAccountServiceContract
	db   *gorm.DB
}

func NewBankAccountController(serv service.BankAccountServiceContract, db *gorm.DB) BankAccountController {
	return BankAccountController{
		serv: serv,
		db:   db,
	}
}

func (obj *BankAccountController) Create(ctx *gin.Context) {
	user, ok := ctx.Request.Context().Value(middleware.UserDataContext).(domain.UserDetail)
	if !ok {
		err := fmt.Errorf("Invalid pasrse data")
		errData := errorhandler.ErrorMapping(1201, "Gagal mendapatakan data - GO-1201", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"response_code": 1201,
			"message":       "Gagal mendapatakan data - GO-1201",
			"data":          errData,
		})
		return
	}

	data := domain.BankAccount{}
	err := ctx.BindJSON(&data)
	if err != nil {
		errData := errorhandler.ErrorMapping(1201, "Gagal mendapatakan data - GO-1201", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"response_code": 1201,
			"message":       "Gagal mendapatakan data - GO-1201",
			"data":          errData,
		})
		return
	}

	data.UserId = user.ID
	errData := obj.serv.Create(obj.db, data)
	if errData.ResponseCode != 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"response_code": errData.ResponseCode,
			"message":       errData.Message,
			"data":          errData,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"response_code": 0,
		"message":       "success",
		"data":          nil,
	})
	return
}
