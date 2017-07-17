package backend

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-redis/redis"
	"os"
)

func NewAwsSession(region string) *session.Session {
	os.Setenv("AWS_SDK_LOAD_CONFIG", "true")
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	return sess
}

func NewRedisSession(host string, port string, pass string, db int) *redis.Client {
	addr := host + ":" + port
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass, // "" for no password
		DB:       db,   // default DB : 0
	})
}
