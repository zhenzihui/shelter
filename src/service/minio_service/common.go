package minio_service

import "projectsuika.com/shelter/src/struct/file_util"

// minio 配置在这里 TODO:从文件读取配置
var config = file_util.MinIOConfig{
	Endpoint:        "127.0.0.1:9000",
	AccessKeyID:     "aFQk1oQuFQm6vixPIyZR",
	SecretAccessKey: "WA1IYZ0mHhs7861znkTC7XQ2zhV8uBfotP5sZFdr",
	UseSSL:          false,
}

// minio singleton客户端
var minioClient = file_util.GetClient(config).MinIO

// 创建文件夹时的隐藏文件
var folderFileName = ".ff"
