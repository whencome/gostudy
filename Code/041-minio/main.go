/**
 * basic toturial:
 *     1. https://zhuanlan.zhihu.com/p/580822213
 *     2. https://blog.csdn.net/weixin_45903371/article/details/114117554
 */
package main

import (
	"context"
	"log"

	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// 定义minio配置
var (
	minioUrl       = "192.168.101.13"
	minioPort      = 9000
	minioEndpoint  = "192.168.101.13:9000"
	minioAccessKey = "admin"
	minioSecretKey = "admin123"
	minioUseSSL    = false
)

func main() {
	// 根据配置连接minio server
	minioClient, err := minio.New(
		minioEndpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(minioAccessKey, minioSecretKey, ""),
			Secure: minioUseSSL,
		})
	if err != nil {
		log.Fatalf("conetct minio server fail %s url %s ", err, minioEndpoint)
	}

	bucketName := "mybucket"
	location := "chengdu"
	// 初始化bucket
	ctx := context.Background()
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, err := minioClient.BucketExists(ctx, bucketName)
		if err == nil && exists {
			log.Printf("owned %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("create bucket %s success\n", bucketName)
	}

	// 上传文件
	localFile := "/mnt/c/Users/eric/Pictures/The-planet-energy-light-space_2560x1440.jpg"
	objectName := "The-planet-energy-light-space_2560x1440.jpg"
	contentType := "image/jpeg"
	n, err := minioClient.FPutObject(ctx, bucketName, objectName, localFile, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalf("put object fail: %s", err)
	}
	log.Printf("put object %s success, result: %#v\n", objectName, n)
}
