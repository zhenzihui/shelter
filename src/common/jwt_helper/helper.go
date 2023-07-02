package jwt_helper

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	_const "projectsuika.com/shelter/src/common/const"
	"time"
)

// SignedDetails 签名内容
type SignedDetails struct {
	UserName string
	UserId   string
	jwt.RegisteredClaims
}

func GenerateAllTokens(userId string, userName string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		UserId:   userId,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(24*3600))),
		},
	}

	refreshClaims := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(24*3600))),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(_const.SecretKey))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(_const.SecretKey))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}
func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(_const.SecretKey), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}
	expireTime, _ := claims.GetExpirationTime()
	if expireTime.Unix() < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}
	return claims, msg
}
