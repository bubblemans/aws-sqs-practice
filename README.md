# aws-sqs-practice

Easy AWS SQS practice of send/receive/delete.

This application is to simulate producer-consumer model using SQS as a middleware. A producer will produce several messages asynchronously, and a worker pool is used to consume messages.

## Usage
```
go run .
```

## TODO
- Clean up codes using interface
- How to optimize the app by changing `MaxNumberOfMessages` in `ReceiveMessageInput`?