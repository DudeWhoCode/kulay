package aws

import (
"github.com/aws/aws-sdk-go/aws/session"
)

var sess *session.Session

func NewSession() (*session.Session) {
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return sess
}
