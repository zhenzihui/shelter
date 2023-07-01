package minio_service

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
)

// Upload 上传到minio
func Upload(fileReader io.Reader, fileName string, fileSize int64, bucket string) minio.UploadInfo {
	info, err := minioClient.PutObject(context.Background(), bucket, fileName, fileReader, fileSize, minio.PutObjectOptions{})
	if err != nil {
		panic(err)
	}
	return info
}
