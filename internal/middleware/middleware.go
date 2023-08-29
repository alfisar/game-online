package middleware

import (
	"context"
	_serv "game-online/application/user/service"
	"game-online/internal/errorhandler"
	"game-online/internal/jwthandler"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AuthenticateMiddleware is a middleware for user authentication
type AuthenticateMiddleware struct {
	serv _serv.UserServiceContract
	jwt  *jwthandler.JwtHandler
	db   *gorm.DB
}

// NewAuthenticateMiddleware create objcet of authenticate middleware
func NewAuthenticateMiddleware(serv _serv.UserServiceContract, jwt *jwthandler.JwtHandler, db *gorm.DB) *AuthenticateMiddleware {
	return &AuthenticateMiddleware{
		serv: serv,
		jwt:  jwt,
		db:   db,
	}
}

type key string

// UserDataContext is a key that used on storing user in context
const UserDataContext key = "user"

// Authenticate authenticates the user who accessed the handler
func (obj *AuthenticateMiddleware) Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, errData := getTokenRequest(ctx.Request)
		if errData.ResponseCode != 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"response_code": 1301,
				"message":       "Failed Get Header Token - GO-1301",
				"data":          nil,
			})
			return

		}

		_, claim, err := obj.jwt.ValidationToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"response_code": 1301,
				"message":       "Invalid data token - GO-1301",
				"data":          nil,
			})
			return
		}
		id, ok := claim["sub"].(float64)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"response_code": 1302,
				"message":       "Something wrong - GO-1302",
				"data":          nil,
			})
			return
		}

		user, errData := obj.serv.GetByID(obj.db, int(id))
		if errData.ResponseCode != 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"response_code": errData.ResponseCode,
				"message":       errData.Message,
				"data":          errData.Data,
			})
			return
		}
		newCtx := context.WithValue(ctx.Request.Context(), UserDataContext, user)
		ctx.Request = ctx.Request.WithContext(newCtx)
		ctx.Next()
		// next.ServeHTTP(ctx.Writer, ctx.Request.WithContext(newCtx))
	}
}

func getTokenRequest(r *http.Request) (tokenString string, err errorhandler.ErrorData) {
	const bearerSchema = "Bearer "
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) < len(bearerSchema) {
		err = errorhandler.ErrorMapping(1301, "please, provide valid authentication - GO-1301", nil)
		return
	}

	tokenString = authHeader[len(bearerSchema):]
	return
}
