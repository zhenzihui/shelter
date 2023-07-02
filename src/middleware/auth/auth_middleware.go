package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_const "projectsuika.com/shelter/src/common/const"
	"projectsuika.com/shelter/src/common/jwt_helper"
	"projectsuika.com/shelter/src/common/response"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(_const.TokenName)
		//没有token
		if token == "" {
			c.JSON(400, response.Error(response.CodeRegistry.NoAuth, ""))
			c.Abort()
			return
		}
		//验证token
		detail, msg := jwt_helper.ValidateToken(token)

		if detail == nil {
			c.JSON(http.StatusBadRequest, response.Error(response.CodeRegistry.InvalidToken, msg))
			c.Abort()
			return
		}
		c.Set(_const.KeyUserId, detail.UserId)
		c.Set(_const.KeyUserName, detail.UserName)
		c.Next()
	}
}
