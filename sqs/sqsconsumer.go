package sqsapp

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"encoding/json"
)
var svc *sqs.SQS

func Consume() {
	sess := NewSession()
	svc = sqs.New(sess)
	qURL := "test_queue"
	fmt.Println("CREATED SVC")
	for {
		result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
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
			var parsed interface{}
			if err := json.Unmarshal([]byte(*msg.Body), &parsed); err != nil {
				fmt.Println("Error parsing json")
			}
			resultDelete, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      &qURL,
				ReceiptHandle: msg.ReceiptHandle,
			})
			if err != nil {
				fmt.Println("Delete Error", err)
				return
			}
			fmt.Println("Message Deleted", resultDelete)
			fmt.Println(parsed)
		}
	}

}