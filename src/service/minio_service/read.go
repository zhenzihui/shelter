package minio_service

import (
	"context"
	"github.com/minio/minio-go/v7"
	"log"
)

// ListObjects bucket下所有文件, path是路径
func ListObjects(bucket string, path string, recursive bool) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	objChan := minioClient.ListObjects(ctx, bucket, minio.ListObjectsOptions{Recursive: recursive, Prefix: path})
	for obj := range objChan {
		log.Println(obj)
	}
}
