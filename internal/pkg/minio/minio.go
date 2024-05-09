package minio

import (
	"bytes"
	"context"
	"dennic_api_gateway/internal/pkg/config"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func UploadToMinio(cfg *config.Config, objectName string, content []byte, contentLength int64) (string, error) {
	minioClient, err := minio.New(cfg.MinioService.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioService.AccessKey, cfg.MinioService.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		return "", err
	}

	found, err := minioClient.BucketExists(context.Background(), cfg.MinioService.BucketName)
	if err != nil {
		return "", err
	}
	if !found {
		err = minioClient.MakeBucket(context.Background(), cfg.MinioService.BucketName, minio.MakeBucketOptions{})
		if err != nil {
			return "", err
		}
		fmt.Println("Bucket created successfully.")
	} else {
		fmt.Println("Bucket already exists.")
	}

	_, err = minioClient.PutObject(context.Background(), cfg.MinioService.BucketName, objectName, bytes.NewReader(content), int64(len(content)), minio.PutObjectOptions{UserMetadata: map[string]string{"x-amz-acl": "private"}})
	if err != nil {
		return "", err
	}

	objectURL := fmt.Sprintf("%s/%s/%s", cfg.MinioService.Endpoint, cfg.MinioService.BucketName, objectName)
	return objectURL, nil
}

// func GetMinioObject(cfg *config.Config, objectName string) (string, error) {
//   minioClient, err := model_minio.New(cfg.MinioService.Endpoint, &model_minio.Options{
//     Creds:  credentials.NewStaticV4(cfg.MinioService.AccessKey, cfg.MinioService.SecretKey, ""),
//     Secure: false,
//   })
//   if err != nil {
//     return "", err
//   }

//   _, err = minioClient.StatObject(context.Background(), cfg.MinioService.BucketName, objectName, model_minio.StatObjectOptions{})
//   if err != nil {
//     return "", err
//   }

//   objectURL := fmt.Sprintf("%s/%s/%s", cfg.MinioService.Endpoint, cfg.MinioService.BucketName, objectName)
//   return objectURL, nil
// }
