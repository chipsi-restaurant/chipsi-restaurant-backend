package usecase

import (
	"chipsiBackend/bootstrap"
	"context"
	"github.com/minio/minio-go/v7"
	"io"
)

type S3Usecase struct {
	Client     *minio.Client
	BucketName string
}

func NewS3Usecase(client *minio.Client, cfg *bootstrap.Config) S3Usecase {
	return S3Usecase{
		Client:     client,
		BucketName: cfg.S3.Bucket,
	}
}

func (s *S3Usecase) UploadImage(ctx context.Context, file io.Reader, objectName, contentType string, size int64) (string, error) {
	_, err := s.Client.PutObject(ctx, s.BucketName, objectName, file, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", err
	}

	url := "https://" + s.BucketName + ".s3.cloud.ru/" + objectName
	return url, nil
}
