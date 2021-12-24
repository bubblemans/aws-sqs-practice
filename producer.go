package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func produce(svc *sqs.SQS, queueURL string) {
	sendMessageInput := sqs.SendMessageInput{
		MessageBody: aws.String("A message from Alvin Lin."),
		QueueUrl:    aws.String(queueURL),
	}

	sendMessageOutput, err := svc.SendMessage(&sendMessageInput)
	checkError(err)

	log.Println(sendMessageOutput)
}
