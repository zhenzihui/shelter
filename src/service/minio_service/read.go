package minio_service

import (
	"context"
	"github.com/minio/minio-go/v7"
	"log"
	"net/url"
	"time"
)

// ListObjects bucket下所有文件, path是路径
func ListObjects(bucketName string, path string, recursive bool) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	objChan := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Recursive: recursive, Prefix: path})
	for obj := range objChan {
		log.Println(obj)
	}
}

func GetDownloadUrl(bucketName string, fileName string) *url.URL {
	reqParams := make(url.Values)
	presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, fileName,
		time.Hour*24*7, reqParams)
	if err != nil {
		panic(err)
	}
	return presignedURL
}
