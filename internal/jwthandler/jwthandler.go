package jwthandler

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtHandler struct {
	Secret string
}

func GetJwt() *JwtHandler {
	return &JwtHandler{}
}

func (obj *JwtHandler) GetToken(lifepan time.Duration, userID int) (token string, err error) {
	exp := time.Now().Add(lifepan)
	claims := jwt.MapClaims{
		"exp": exp.Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"sub": userID,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = jwtToken.SignedString([]byte(obj.Secret))
	if err != nil {
		return
	}

	return
}

func (obj *JwtHandler) ValidationToken(token string) (isValid bool, claim jwt.MapClaims, err error) {
	jwttoken, err := jwt.Parse(token, func(jwttoken *jwt.Token) (interface{}, error) {
		if _, ok := jwttoken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		obj.Secret = os.Getenv("JWT_SECRET")
		return []byte(obj.Secret), nil
	})

	if err != nil {
		return
	}

	if !jwttoken.Valid {
		return
	}

	claim, ok := jwttoken.Claims.(jwt.MapClaims)
	if !ok {
		return
	}

	isValid = true
	return

}
