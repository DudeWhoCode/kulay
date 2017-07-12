package backend

import (
	"reflect"
	"testing"
)

func TestNewAwsSession(t *testing.T) {
	region := "us-east-1"
	sess := NewAwsSession(region)
	if sessType := reflect.TypeOf(sess).String(); sessType == "*session.Session" {
		t.Log("Received expected session type")
	} else {
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
	client := NewRedisSession(host, port, pass, db)
	if pong, err := client.Ping().Result(); pong != "PONG" {
		t.Errorf("Expected PONG, got %v", pong)
		t.Errorf("Expected no errors, got %s", err)
	}
}