package controller

import (
	"fmt"
	"game-online/application/wallet/service"
	"game-online/domain"
	"game-online/internal/errorhandler"
	"game-online/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WalletController struct {
	serv service.WalletLogServiceContract
	db   *gorm.DB
}

type keyCtx string

const (
	IDCTXKey     keyCtx = "user_id"
	WalletCTXKey keyCtx = "wallet_id"
)

func NewWalletController(serv service.WalletLogServiceContract, db *gorm.DB) WalletController {
	return WalletController{
		serv: serv,
		db:   db,
	}
}

func (obj *WalletController) TopUp(ctx *gin.Context) {
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

	req := domain.TopUp{}
	err := ctx.BindJSON(&req)
	if err != nil {
		errData := errorhandler.ErrorMapping(1201, "Gagal mendapatakan data - GO-1201", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"response_code": 1201,
			"message":       "Gagal mendapatakan data - GO-1201",
			"data":          errData,
		})
		return
	} else if req.Nominal <= 0 {
		err := fmt.Errorf("Invalid data request")
		errData := errorhandler.ErrorMapping(1202, "Invalid data request - GO-1202", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"response_code": 1201,
			"message":       "Invalid data request - GO-1202",
			"data":          errData,
		})
		return
	}

	errData := obj.serv.Transaction(obj.db, req.Nominal, user.ID, user.WalletID)
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
