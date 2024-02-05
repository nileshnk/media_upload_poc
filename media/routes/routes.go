package routes

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/go-chi/chi/v5"
	Config "github.com/nileshnk/media_upload_poc/media/config"
)

func MainRouter(ctx context.Context) chi.Router {

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Media Service"))
	})

	r.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		r.ParseForm()
		file, fileHeader, err := r.FormFile("file")
		CheckErrorWithPanic(err)
		fileBytes := make([]byte, fileHeader.Size)
		n, errRead := file.Read(fileBytes)
		CheckErrorWithPanic(errRead)
		if n != int(fileHeader.Size) {
			fmt.Println(n);
		}
		
		cwd, errCwd := os.Getwd()
		CheckErrorWithPanic(errCwd)

		newFile, errCreate := os.Create(cwd + "/files/" + fileHeader.Filename)
		CheckErrorWithPanic(errCreate)
		defer newFile.Close()	
		S3Upload(ctx, file, fileHeader.Filename)

		_, errWrite := newFile.Write(fileBytes);
		CheckErrorWithPanic(errWrite)


		w.Write([]byte("File Uploaded"))
	})

	return r
}
func CheckErrorWithPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func S3Upload (ctx context.Context,fileBytes io.Reader, fileName string) {
	AwsConfig := GetS3Config()
	AwsSession := GetS3Session(&AwsConfig)
	uploader := s3manager.NewUploader(AwsSession);
	_, err := uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket: &Config.GetConfig.Media.AWS.BucketName,
		Key: &fileName,
		Body: fileBytes,
	})
	CheckErrorWithPanic(err)
}

func GetS3Config() aws.Config {
	KeyID := Config.GetConfig.Media.AWS.AccessKeyID
	KeySecret := Config.GetConfig.Media.AWS.SecretAccessKey
	EndPoint := Config.GetConfig.Media.AWS.EndPoint
	var config aws.Config = aws.Config{
		CredentialsChainVerboseErrors: aws.Bool(true),
		Region: &Config.GetConfig.Media.AWS.Region,
		Endpoint: &EndPoint, //&endPointUrl, //
		DisableSSL: aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
		Credentials: credentials.NewStaticCredentials(KeyID, KeySecret, ""),
	}
	return config
}

func GetS3Session(config *aws.Config) *session.Session {
	return session.Must(session.NewSession(
		config,
	))
}