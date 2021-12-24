package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/joho/godotenv"
)

const (
	MESSAGE_COUNT = 10
	NUM_WORKERS   = 3
)

func main() {
	svc, queueURL := setup()

	// Publish messages to SQS asynchronously
	for i := 0; i < MESSAGE_COUNT; i++ {
		go produce(svc, queueURL)
	}

	// Use a worker pool to consume messages
	consume(svc, queueURL)
}

func setup() (*sqs.SQS, string) {
	err := godotenv.Load()
	checkError(err)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION"))},
	)
	checkError(err)

	svc := sqs.New(sess)
	queueURL := os.Getenv("QUEUE_URL")
	return svc, queueURL
}
