package controller

import (
	"fmt"
	"game-online/application/user/service"
	"game-online/domain"
	"game-online/internal/errorhandler"
	"game-online/internal/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	serv service.UserServiceContract
	db   *gorm.DB
}

func NewUserController(serv service.UserServiceContract, db *gorm.DB) UserController {
	return UserController{
		serv: serv,
		db:   db,
	}
}

func (obj *UserController) Register(ctx *gin.Context) {
	data := domain.User{}
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

func (obj *UserController) Login(ctx *gin.Context) {
	data := domain.User{}
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

	token, errData := obj.serv.Login(obj.db, data.Name, data.Password)
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
		"data":          token,
	})
	return
}

func (obj *UserController) Logout(ctx *gin.Context) {
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

	errData := obj.serv.Logout(obj.db, user.ID)
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

func (obj UserController) GetList(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Request.URL.Query().Get("page"))
	if err != nil {
		errData := errorhandler.ErrorMapping(1201, "Gagal mendapatakan data param - GO-1203", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"response_code": errData.ResponseCode,
			"message":       errData.Message,
			"data":          errData,
		})
		return
	}
	itemPages, err := strconv.Atoi(ctx.Request.URL.Query().Get("item_per_page"))
	if err != nil {
		errData := errorhandler.ErrorMapping(1201, "Gagal mendapatakan data param - GO-1203", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"response_code": errData.ResponseCode,
			"message":       errData.Message,
			"data":          errData,
		})
		return
	}
	search := ctx.Request.URL.Query().Get("search")
	_, dataPagination, errData := obj.serv.GetList(obj.db, page, itemPages, search)
	if errData.ResponseCode != 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"response_code": errData.ResponseCode,
			"message":       errData.Message,
			"data":          errData,
		})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"response_code": 0,
		"message":       "success",
		"data":          dataPagination,
	})
	return
}

func (obj *UserController) Detail(ctx *gin.Context) {
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

	ctx.JSON(200, gin.H{
		"response_code": 0,
		"message":       "success",
		"data":          user,
	})
	return
}
