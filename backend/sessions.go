package backend

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
)

var sess *session.Session

func NewAwsSession(region string) *session.Session {
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(region)},
	}))
	return sess
}
