package backend

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"os"
)

func NewAwsSession(region string) *session.Session {
	os.Setenv("AWS_SDK_LOAD_CONFIG", "true")
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	return sess
}
