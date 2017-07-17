package redispubsub

import (
	"github.com/DudeWhoCode/kulay/backend"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	host := "localhost"
	port := "6379"
	pass := ""
	db := 0
	channel := "test"
	testCnt := 5
	pipe := make(chan string, testCnt)
	client := backend.NewRedisSession(host, port, pass, db)
	testMsg := `{ "name": "kulay",
				  "desc"; ""High speed message routing between services",
				  "https://github.com/dudewhocode/kulay",
				  135
				  }`
	go Get(host, port, pass, db, channel, pipe)
	time.Sleep(500 * time.Millisecond)
	for i := 1; i <= testCnt; i++ {
		if err := client.Publish(channel, testMsg).Err(); err != nil {
			t.Errorf("Expected no errors while sending test message to redis, got %s", err)
		}
	}
	// unsubscribe from channel
	if err := client.Publish(channel, "$^KILL^$").Err(); err != nil {
		t.Errorf("Expected no errors while sending test message to redis, got %s", err)
	}
	time.Sleep(500 * time.Millisecond)
	if len(pipe) != testCnt {
		t.Errorf("Expected message count in channel is %v, got %v", testCnt, len(pipe))
	}

}

func TestPut(t *testing.T) {
	host := "localhost"
	port := "6379"
	pass := ""
	db := 0
	channel := "test"
	testCnt := 5
	pipe := make(chan string, testCnt)
	client := backend.NewRedisSession(host, port, pass, db)
	testMsg := `{ "name": "kulay",
				  "desc"; ""High speed message routing between services",
				  "https://github.com/dudewhocode/kulay",
				  135
				  }`
	var testResults []string
	go Put(host, port, pass, db, channel, pipe)
	pubsub := client.Subscribe(channel)
	time.Sleep(time.Second)
	go func() {
		for {
			msg, err := pubsub.ReceiveMessage()
			if err != nil {
				t.Errorf("Expected no errors while receiving message from pubsub channel, got %s", err)
			}
			if msg.Payload == "$^KILL^$" {
				break
			}
			testResults = append(testResults, msg.Payload)
		}
	}()
	for i := 1; i <= testCnt; i++ {
		pipe <- testMsg
	}
	pipe <- "$^KILL^$"
	time.Sleep(time.Second)
	if len(testResults) != testCnt {
		t.Errorf("Expected received message count from redis pubsub is %v, got %v", testCnt, len(testResults))
	}
}