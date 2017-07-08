package sqsapp

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
	. "github.com/DudeWhoCode/kulay/logger"
	"github.com/DudeWhoCode/kulay/backend"
)

var svc *sqs.SQS

func Put(qURL string, region string, rec <-chan string) {
	sess := backend.NewAwsSession(region)
	svc = sqs.New(sess)
	for msg := range rec {
		result, err := svc.SendMessage(&sqs.SendMessageInput{
			DelaySeconds: aws.Int64(10),
			MessageBody:  aws.String(msg),
			QueueUrl:     &qURL,
		})
		if err != nil {
			Log.Error("Error while sending message : ", err)
			continue
		}
		Log.Info("Sent message to SQS queue : ", *result.MessageId)
	}
}


func Get(qURL string, region string, del bool, snd chan<- string) {
	sess := backend.NewAwsSession(region)
	svc = sqs.New(sess)
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
		Log.Info("Got 10 msgs from sqs")
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
				_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
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
		Log.Info("EXITED FOR LOOP")
	}
	Log.Info("EXITED FOREVER LOOP")

}


