package aws

import (
	"bytes"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/estenssoros/sheetdrop/constants"
	"github.com/pkg/errors"
	"github.com/seaspancode/theseus/helpers"
)

// S3SaveFile saves a file to s3
func S3SaveFile(fileName string, bs []byte) error {
	sess, err := newSession()
	if err != nil {
		return errors.Wrap(err, "newSession")
	}
	_, err = s3manager.NewUploader(sess).Upload(&s3manager.UploadInput{
		Bucket: aws.String(constants.SheetDropBucket),
		Key:    aws.String(fileName),
		Body:   bytes.NewBuffer(bs),
	})
	return errors.Wrap(err, "Upload")
}

// S3DownloadFile downloads a file from s3
func S3DownloadFile(fileName string) ([]byte, error) {
	sess, err := newSession()
	if err != nil {
		return nil, errors.Wrap(err, "newSession")
	}
	buffer := &aws.WriteAtBuffer{}
	_, err = s3manager.NewDownloader(sess).Download(buffer, &s3.GetObjectInput{
		Bucket: aws.String(constants.SheetDropBucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return nil, errors.Wrap(err, "downloader.Download")
	}
	return buffer.Bytes(), nil
}

// S3FileExists quick head method to see if file exists
func S3FileExists(fileName string) (bool, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		return false, errors.Wrap(err, "new session")
	}
	_, err = s3.New(sess).HeadObject(&s3.HeadObjectInput{
		Bucket: helpers.StringPtr(constants.SheetDropBucket),
		Key:    helpers.StringPtr(fileName),
	})
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return false, nil
		}
		return false, errors.Wrap(err, "svc.HeadObject")
	}
	return true, nil
}

func newSession() (*session.Session, error) {
	id := os.Getenv(constants.SheetDropAWSAccessKeyID)
	secret := os.Getenv(constants.SheetDropAWSSecretAccessKey)
	if id == "" || secret == "" {
		return session.NewSession(&aws.Config{
			Region: aws.String(constants.AWSRegion),
		})
	}
	return session.NewSession(&aws.Config{
		Region:      aws.String(constants.AWSRegion),
		Credentials: credentials.NewStaticCredentials(id, secret, ""),
	})
}
