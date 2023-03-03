package awsS3handle

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	S3 *s3.Client
}

// This function initializes the AWS S3 client object and returns it.
func InitS3Client() S3Client {
	// Get AWS Configurations.
	sdkConfig := getAWSConfig()

	// Create an Amazon S3 service client
	s3Client := s3.NewFromConfig(sdkConfig)

	return S3Client{S3: s3Client}
}

// This function returns the bucket object content in stream of bytes.
func (s S3Client) GetObjectContent(bucketName string, objectKey string) []byte {
	// Get the object content.
	data, err := s.S3.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})

	checkError(err)

	// Close the content body
	defer data.Body.Close()

	// Get the content body.
	content, e := io.ReadAll(data.Body)
	checkError(e)

	return content
}

// This Function uploads the file into bucket and object specified.
func (s S3Client) UploadObject(bucketName string, objectKey string, filePath string) {
	// Get the file object.
	file, e := os.Open(filePath)
	checkError(e)

	// Upload the file into the bucket.
	_, err := s.S3.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	checkError(err)

	// Console Message.
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println("Object Upload Completed..")
	fmt.Printf("\nBucket Name - %s \nObject - %s \nFileName - %s\n", bucketName, objectKey, filePath)
	fmt.Println(strings.Repeat("-", 100))

	// Close the file.
	file.Close()
}

// This function gets the AWS system configuration.
func getAWSConfig() aws.Config {
	// Load the Shared AWS Configuration (~/.aws/config)
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	checkError(err)
	return sdkConfig
}

// This function checks the error.
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
