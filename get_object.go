package sikalabs_s3_go

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetObject(cfg S3Config, name string) ([]byte, error) {
	ctx := context.TODO()

	awscfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.AccessKey,
				cfg.SecretKey,
				""),
		),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(awscfg)

	result, err := client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &cfg.BucketName,
		Key:    &name,
	})
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	content, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}
