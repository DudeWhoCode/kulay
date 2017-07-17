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
	testMsg := `{ "name": "kulay",
				  "desc"; ""High speed message routing between services",
				  "https://github.com/dudewhocode/kulay",
				  135
				  }`
	go Get(host, port, pass, db, channel, pipe)
	time.Sleep(500 * time.Millisecond)
	client := backend.NewRedisSession(host, port, pass, db)
	for i := 1; i <= testCnt; i++ {
		if err := client.Publish(channel, testMsg).Err(); err != nil {
			panic(err)
		}
	}
	// unsubscribe from channel
	if err := client.Publish(channel, "$^KILL^$").Err(); err != nil {
		panic(err)
	}
	time.Sleep(500 * time.Millisecond)
	if len(pipe) != testCnt {
		t.Errorf("Expected message count in channel is %v, got %v", testCnt, len(pipe))
	}

}
