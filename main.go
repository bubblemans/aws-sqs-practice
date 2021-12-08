package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	checkError(err)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	checkError(err)

	svc := sqs.New(sess)
	queueURL := os.Getenv("QUEUE_URL")

	// send(svc, queueURL)
	msgResult := receive(svc, queueURL)
	delete(svc, queueURL, msgResult)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func send(svc *sqs.SQS, queueURL string) {
	sendMessageInput := sqs.SendMessageInput{
		MessageBody: aws.String("A message from Alvin Lin."),
		QueueUrl:    aws.String(queueURL),
	}

	sendMessageOutput, err := svc.SendMessage(&sendMessageInput)

	checkError(err)

	log.Println(sendMessageOutput)
}

func receive(svc *sqs.SQS, queueURL string) *sqs.ReceiveMessageOutput {
	msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: aws.String(queueURL),
	})
	checkError(err)
	log.Println(msgResult)
	return msgResult
}

func delete(svc *sqs.SQS, queueURL string, msgResult *sqs.ReceiveMessageOutput) {
	for _, message := range msgResult.Messages {
		_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      aws.String(queueURL),
			ReceiptHandle: message.ReceiptHandle,
		})
		checkError(err)
	}
}
