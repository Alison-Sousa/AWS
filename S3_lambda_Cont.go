package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	Name   string
	Gender string
}

func processCSVContent(csvContent io.Reader) ([]Student, error) {
	var students []Student

	csvReader := csv.NewReader(csvContent)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if len(record) != 2 {
			return nil, fmt.Errorf("unexpected number of fields in CSV record: %v", record)
		}
		students = append(students, Student{Name: record[0], Gender: record[1]})
	}

	return students, nil
}

func calculateStatistics(students []Student) (int, int) {
	maleCount := 0
	femaleCount := 0
	for _, student := range students {
		if student.Gender == "M" {
			maleCount++
		} else if student.Gender == "F" {
			femaleCount++
		}
	}
	return maleCount, femaleCount
}

func storeStatisticsInDatabase(maleCount, femaleCount int) error {
	db, err := sql.Open("mysql", "user:password@tcp(database-host:3306)/database-name")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO statistics (timestamp, male_count, female_count) VALUES (?, ?, ?)", time.Now(), maleCount, femaleCount)
	if err != nil {
		return err
	}

	return nil
}

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

		students, err := processCSVContent(obj.Body)
		if err != nil {
			log.Printf("Error processing CSV content from object %s in bucket %s: %v", s3Object, s3Bucket, err)
			continue
		}

		maleCount, femaleCount := calculateStatistics(students)

		err = storeStatisticsInDatabase(maleCount, femaleCount)
		if err != nil {
			log.Printf("Error storing statistics in database: %v", err)
			continue
		}

		log.Printf("Processed CSV file %s in bucket %s. Male count: %d, Female count: %d", s3Object, s3Bucket, maleCount, femaleCount)
	}

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
