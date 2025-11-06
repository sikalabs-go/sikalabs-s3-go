package sikalabs_s3_go

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Config struct {
	BucketName string
	Region     string
	AccessKey  string
	SecretKey  string
}

func PutObject(cfg S3Config, name string, content []byte) error {
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
		return err
	}

	client := s3.NewFromConfig(awscfg)

	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(cfg.BucketName),
		Key:    aws.String(name),
		Body:   bytes.NewReader(content),
	})
	if err != nil {
		return err
	}

	return nil
}
