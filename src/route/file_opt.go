package route

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectsuika.com/shelter/src/service/minio_service"
)

var path = "/files"

func InitFileOperateRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST(path+"/upload/single", func(c *gin.Context) {
		fileHeader, _ := c.FormFile("file")
		file, _ := fileHeader.Open()
		defer file.Close()
		reader := bufio.NewReader(file)
		minio_service.Upload(reader, fileHeader.Filename, fileHeader.Size, "test1")

		c.JSON(http.StatusOK, fmt.Sprintf("'%s' uploaded!", fileHeader.Filename))
	})
	r.GET(path+"/list-mine", func(c *gin.Context) {
		minio_service.ListObjects("test1", "", false)
		c.JSON(http.StatusOK, "ok")

	})

}
