package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func handleRequest(ctx context.Context, s3Event events.S3Event) error {
	sess := session.Must(session.NewSession())

	for _, record := range s3Event.Records {
		s3Bucket := record.S3.Bucket.Name
		s3Object := record.S3.Object.Key

		s3Svc := s3.New(sess)
		obj, err := s3Svc.GetObject(&s3.GetObjectInput{Bucket: &s3Bucket, Key: &s3Object})
		if err != nil {
			log.Printf("Error getting object %s from bucket %s: %v", s3Object, s3Bucket, err)
			continue
		}

		buf := make([]byte, 1024)
		n, err := obj.Body.Read(buf)
		if err != nil {
			log.Printf("Error reading object %s from bucket %s: %v", s3Object, s3Bucket, err)
			continue
		}

		fmt.Printf("Content of object %s in bucket %s:\n%s\n", s3Object, s3Bucket, string(buf[:n]))
	}

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
