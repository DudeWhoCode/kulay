package backend

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

var sess *session.Session

func NewAwsSession() *session.Session {
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return sess
}
