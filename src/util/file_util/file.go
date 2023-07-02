package file_util

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
	"sync"
)

// MinIOConfig minio配置
type MinIOConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

type Client struct {
	MinIO *minio.Client
}

var lock = &sync.Mutex{}
var singleClient *Client
var UserHome, _ = os.UserHomeDir()

// GetClient 获取单例client
func GetClient(config MinIOConfig) *Client {
	if singleClient == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleClient == nil {
			singleClient = &Client{
				MinIO: GetMinIOClient(config),
			}
		} else {
			log.Println("already has client")
		}
	} else {
		log.Println("already created client")
	}
	return singleClient
}

// GetMinIOClient 直接得到MinIO对象
func GetMinIOClient(config MinIOConfig) *minio.Client {
	client, err := minio.New(config.Endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
			Secure: config.UseSSL,
		},
	)
	if err != nil {
		log.Printf("error with initializing minio_service client: %s", err)
		return nil
	}
	return client
}
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
