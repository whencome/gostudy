/**
 * basic toturial:
 *     1. https://zhuanlan.zhihu.com/p/580822213
 *     2. https://blog.csdn.net/weixin_45903371/article/details/114117554
 */
package main

import (
    "context"
    "fmt"
    "log"

    minio "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
)

// 定义minio配置
var (
    minioEndpoint  = "127.0.0.1:9000"
    minioAccessKey = "admin"
    minioSecretKey = "admin123"
    minioUseSSL    = false
)

func getPublicPolicy(bucketName string) string {
    // version只能写死为2012-10-17
    return `{
	"Version": "2012-10-17",
	"Statement": [{
		"Effect": "Allow",
		"Principal": {
			"AWS": ["*"]
		},
		"Action": ["s3:GetBucketLocation", "s3:ListBucket", "s3:ListBucketMultipartUploads"],
		"Resource": ["arn:aws:s3:::` + bucketName + `"]
	}, {
		"Effect": "Allow",
		"Principal": {
			"AWS": ["*"]
		},
		"Action": ["s3:AbortMultipartUpload", "s3:DeleteObject", "s3:GetObject", "s3:ListMultipartUploadParts", "s3:PutObject"],
		"Resource": ["arn:aws:s3:::` + bucketName + `/*"]
	}]
}`
}

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

    bucketName := "mybucket01"
    location := "cn-chengdu-01"
    policy := getPublicPolicy(bucketName)
    // 初始化bucket
    ctx := context.Background()
    // 1. 先检查bucket是否存在
    exists, err := minioClient.BucketExists(ctx, bucketName)
    if err == nil && !exists {
        log.Printf("bucket %s not exists, try to make it\n", bucketName)
        err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
        if err != nil {
            log.Panicf("make bucket err: %s\n", err)
        }
        // 设置策略(设置为公共策略方便直接通过地址访问)
        err = minioClient.SetBucketPolicy(ctx, bucketName, policy)
        if err != nil {
            log.Fatalln("set policy err: ", err)
        }
    } else {
        if err != nil {
            log.Panicf("check bucket err: %s\n", err)
        } else {
            log.Printf("bucket %s exists", bucketName)
        }
    }

    // 上传文件
    localFile := "testimg.jpg"
    objectName := "2023/05/25/testimg.jpg"
    contentType := "image/jpeg"
    n, err := minioClient.FPutObject(ctx, bucketName, objectName, localFile, minio.PutObjectOptions{ContentType: contentType})
    if err != nil {
        log.Fatalf("put object fail: %s", err)
    }
    log.Printf("put object %s success, result: %#v\n", objectName, n)

    // 构造访问地址
    imgUrl := fmt.Sprintf("%s/%s/%s", minioEndpoint, bucketName, objectName)
    log.Printf("image url: %#v\n", imgUrl)
}
