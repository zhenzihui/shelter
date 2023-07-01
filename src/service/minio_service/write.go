package minio_service

import (
	"context"
	"fmt"
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

// Delete 删除文件
func Delete(bucket string, fileName string) error {
	err := minioClient.RemoveObject(context.Background(), bucket, fileName, minio.RemoveObjectOptions{})
	return err
}

func NewFolder(bucketName string, path string) minio.UploadInfo {
	info, err := minioClient.PutObject(context.Background(), bucketName,
		fmt.Sprintf("%s/%s", path, folderFileName),
		nil, 0, minio.PutObjectOptions{})
	if err != nil {
		panic(err)
	}
	return info
}
