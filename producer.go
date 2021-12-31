package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func (client Client) produce() {
	sendMessageInput := sqs.SendMessageInput{
		MessageBody: aws.String("A message from Alvin Lin."),
		QueueUrl:    aws.String(client.queueURL),
	}

	sendMessageOutput, err := client.svc.SendMessage(&sendMessageInput)
	checkError(err)

	log.Println(sendMessageOutput)
}
