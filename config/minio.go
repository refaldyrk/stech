package config

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"log"
	"time"
)

func (c *Config) minioConnect() {
	timeStart := time.Now()
	endpoint := viper.Get("MINIO_ENDPOINT").(string)
	accessKeyID := viper.Get("MINIO_ACCESS").(string)
	secretAccessKey := viper.Get("MINIO_SECRET").(string)

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		panic(err.Error())
	}

	minioClient.MakeBucket(context.Background(), viper.GetString("MINIO_KYC_BUCKET"), minio.MakeBucketOptions{})

	log.Println("Minio client created in ", time.Since(timeStart).Milliseconds(), "ms")

	c.MinioClient = minioClient
}
