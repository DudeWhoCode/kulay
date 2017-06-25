package sqsapp

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"naren/kulay/config"
)

var producerSvc *sqs.SQS

func produce(qURL string, rec <-chan string, done chan bool) {
	sess := NewSession()
	producerSvc = sqs.New(sess)
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

func Push(pipe <-chan string, done chan bool, cfg config.Kulay) {
	qURL := cfg.QueueUrl
	fmt.Println("starting go Produce routine")
	produce(qURL, pipe, done)
}
