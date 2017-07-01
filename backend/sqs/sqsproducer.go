package sqsapp

import (
	"naren/kulay/backend"
	. "naren/kulay/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var producerSvc *sqs.SQS

func Put(qURL string, rec <-chan string) {
	sess := backend.NewAwsSession()
	producerSvc = sqs.New(sess)
	for msg := range rec {
		result, err := producerSvc.SendMessage(&sqs.SendMessageInput{
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
