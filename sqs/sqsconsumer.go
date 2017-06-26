package sqsapp

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"naren/kulay/config"
	. "naren/kulay/logger"
)

var consumerSvc *sqs.SQS

func consume(qURL string, snd chan<- string, done chan bool, del bool) {
	sess := NewSession()
	consumerSvc = sqs.New(sess)
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
			Log.Error("Error", err)
			return
		}

		if len(result.Messages) == 0 {
			Log.Warn("Received no messages")
			return
		}
		for _, msg := range result.Messages {
			parsed := msg.Body
			if del == true {
				_, err := consumerSvc.DeleteMessage(&sqs.DeleteMessageInput{
					QueueUrl:      &qURL,
					ReceiptHandle: msg.ReceiptHandle,
				})
				if err != nil {
					Log.Error("Delete Error", err)
					return
				}
				Log.Info("Message Deleted")
			}
			Log.Info("Message Received and sent to channel")
			snd <- *parsed
		}
	}
	done <- true

}

func Get(pipe chan<- string, done chan bool, cfg config.Kulay) {
	qURL := cfg.QueueUrl
	del := cfg.Delete
	go consume(qURL, pipe, done, del)
}
