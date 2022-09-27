package middlewares

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"net/http"
	"os"
)

const (
	AWS_S3_REGION = "us-east-1"
	AWS_S3_BUCKET = "lratajczakmybucketforquestions"
)

func UploadFile(uploadFileDir string) error {
	session, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION), DisableRestProtocolURICleaning: aws.Bool(true)})

	if err != nil {
		log.Fatal(err)
	}
	upFile, err := os.Open(uploadFileDir)
	if err != nil {
		return err
	}
	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
	var fileSize int64 = upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	fileType := http.DetectContentType(fileBuffer)
	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(AWS_S3_BUCKET),
		Key:                  aws.String(uploadFileDir),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(fileBuffer),
		ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String(fileType),
		ServerSideEncryption: aws.String("AES256"),
	})

	return err
}
