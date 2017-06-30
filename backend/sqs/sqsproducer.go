package sqsapp

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"naren/kulay/config"
	. "naren/kulay/logger"
	"naren/kulay/backend"
)

var producerSvc *sqs.SQS

func produce(qURL string, rec <-chan string) {
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

func Put(pipe <-chan string, cfg interface{}) {
	sqsCfg := cfg.(config.SQSConf)
	qURL := sqsCfg.QueueUrl
	produce(qURL, pipe)
}
