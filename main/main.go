package main

import (
	"github.com/gin-gonic/gin"
	"projectsuika.com/shelter/src/route"
)

func main() {
	r := gin.Default()
	route.InitFileOperateRoutes(r)
	r.Run("0.0.0.0:9008")
}
