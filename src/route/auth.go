package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectsuika.com/shelter/src/common/request"
	"projectsuika.com/shelter/src/common/response"
	"projectsuika.com/shelter/src/service/user_service"
)

var authPath = "/user"

func AuthRoutes(r *gin.Engine) {
	r.POST(fmt.Sprintf("%s/login", authPath), func(c *gin.Context) {
		var loginReq request.LoginReq
		err := c.BindJSON(&loginReq)
		if err != nil {
			panic(err)
		}
		resp, err := user_service.Login(loginReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.CodeRegistry.InvalidPassword)
			return
		}
		c.JSON(http.StatusOK, resp)
	})

}
