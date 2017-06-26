package sqsapp

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"naren/kulay/config"
	. "naren/kulay/logger"
)

var producerSvc *sqs.SQS

func produce(qURL string, rec <-chan string, done chan bool) {
	sess := NewSession()
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
	done <- true
}

func Put(pipe <-chan string, done chan bool, cfg interface{}) {
	Log.Println("INTERFACE ", cfg)
	sqsCfg := cfg.(config.SQSConf)
	qURL := sqsCfg.QueueUrl
	produce(qURL, pipe, done)
}
