package sqsapp

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"naren/kulay/config"
)

var consumerSvc *sqs.SQS

func pull(snd chan<- string, done chan bool) {
	sess := NewSession()
	consumerSvc = sqs.New(sess)
	qURL := config.KulayConf.QueueUrl
	for {
		result, err := consumerSvc.ReceiveMessage(&sqs.ReceiveMessageInput{
			AttributeNames: []*string{
				aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
			},
			MessageAttributeNames: []*string{
				aws.String(sqs.QueueAttributeNameAll),
			},
			QueueUrl:            &qURL,
			MaxNumberOfMessages: aws.Int64(10),
			VisibilityTimeout:   aws.Int64(30),
			WaitTimeSeconds:     aws.Int64(20),
		})
		if err != nil {
			fmt.Println("Error", err)
			return
		}

		if len(result.Messages) == 0 {
			fmt.Println("Received no messages")
			return
		}
		for _, msg := range result.Messages {
			parsed := msg.Body
			_, err := consumerSvc.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      &qURL,
				ReceiptHandle: msg.ReceiptHandle,
			})
			if err != nil {
				fmt.Println("Delete Error", err)
				return
			}
			fmt.Println("Message Deleted")
			snd <- *parsed
		}
	}
	done <- true

