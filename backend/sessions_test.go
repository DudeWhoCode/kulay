package backend

import (
	"reflect"
	"testing"
)

func TestNewAwsSession(t *testing.T) {
	region := "us-east-1"
	sess := NewAwsSession(region)
	if sessType := reflect.TypeOf(sess).String(); sessType != "*session.Session" {
		t.Errorf("Expected type *session.Session, got %v", sessType)
	}
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected no panics, got %v", r)
		}
	}()
}

func TestNewRedisSession(t *testing.T) {
	host := "localhost"
	port := "6379"
	pass := ""
	db := 0
	sess := NewRedisSession(host, port, pass, db)
	if sessType := reflect.TypeOf(sess).String(); sessType != "*redis.Client" {
		t.Errorf("Expected type *Client, got %v", sessType)
	}

}
