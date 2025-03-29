package bootstrap

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewS3Client(cfg *Config) (*minio.Client, error) {
	var err error
	s3Client, err := minio.New(cfg.S3.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.S3.AccessKey, cfg.S3.SecretKey, ""),
		Secure: true,
		Region: cfg.S3.Region,
	})
	if err != nil {
		return nil, err
	}
	return s3Client, nil
}
