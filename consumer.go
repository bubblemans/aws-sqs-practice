package main

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func (client Client) consume() {
	// Use a worker pool to consume messages
	messages := make(chan *sqs.ReceiveMessageOutput)

	// spawning workers
	for workerId := 1; workerId <= NUM_WORKERS; workerId++ {
		go worker(client.svc, client.queueURL, workerId, messages)
	}

	for {
		messages <- receive(client.svc, client.queueURL)
	}
}

func worker(svc *sqs.SQS, queueURL string, id int, messages <-chan *sqs.ReceiveMessageOutput) {
	for message := range messages {
		if len(message.Messages) != 0 {
			log.Printf("Worker %d process message", id)
			time.Sleep(1 * time.Second)
			delete(svc, queueURL, message)
		} else {
			log.Println("Empty message")
		}
	}
}

func receive(svc *sqs.SQS, queueURL string) *sqs.ReceiveMessageOutput {
	waitTimeSeconds := 1
	msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:        aws.String(queueURL),
		WaitTimeSeconds: aws.Int64(int64(waitTimeSeconds)), // long polling
	})
	checkError(err)
	log.Println(msgResult)
	return msgResult
}

// Using SQS, we have to delete message from the queue ourselves
func delete(svc *sqs.SQS, queueURL string, msgResult *sqs.ReceiveMessageOutput) {
	for _, message := range msgResult.Messages {
		_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      aws.String(queueURL),
			ReceiptHandle: message.ReceiptHandle,
		})
		checkError(err)
	}
}
