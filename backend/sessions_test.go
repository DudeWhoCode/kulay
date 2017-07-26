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
