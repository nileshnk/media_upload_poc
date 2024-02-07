package services

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	Config "github.com/nileshnk/media_upload_poc/media/config"
)

func S3Upload(ctx context.Context, fileBytes io.Reader, fileName string) error {
	AwsConfig := GetS3Config()
	AwsSession := GetS3Session(&AwsConfig)
	uploader := s3manager.NewUploader(AwsSession)
	_, err := uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket: &Config.GetConfig.Media.AWS.BucketName,
		Key:    &fileName,
		Body:   fileBytes,
	})
	return err
}

func GetS3Config() aws.Config {
	KeyID := Config.GetConfig.Media.AWS.AccessKeyID
	KeySecret := Config.GetConfig.Media.AWS.SecretAccessKey
	EndPoint := Config.GetConfig.Media.AWS.EndPoint
	var config aws.Config = aws.Config{
		CredentialsChainVerboseErrors: aws.Bool(true),
		Region:                        &Config.GetConfig.Media.AWS.Region,
		Endpoint:                      &EndPoint, //&endPointUrl, //
		DisableSSL:                    aws.Bool(true),
		S3ForcePathStyle:              aws.Bool(true),
		Credentials:                   credentials.NewStaticCredentials(KeyID, KeySecret, ""),
	}
	return config
}

func GetS3Session(config *aws.Config) *session.Session {
	return session.Must(session.NewSession(
		config,
	))
}
