package sqsapp

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var producerSvc *sqs.SQS

func produce(rec <-chan string, done chan bool) {
	sess := NewSession()
	producerSvc = sqs.New(sess)
	qURL := "https://sqs.ap-southeast-1.amazonaws.com/971037846030/test"
	fmt.Println("before send message")
	for msg := range rec {
		result, err := producerSvc.SendMessage(&sqs.SendMessageInput{
			DelaySeconds: aws.Int64(10),
			MessageBody:  aws.String(msg),
			QueueUrl:     &qURL,
		})
		if err != nil {
			fmt.Println("Error", err)
			continue
		}
		fmt.Println("Success", *result.MessageId)
	}
	done <- true
}

func Push(pipe <-chan string, done chan bool) {
	fmt.Println("starting go Produce routine")
	produce(pipe, done)
}
