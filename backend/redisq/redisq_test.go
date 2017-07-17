package redisq

import (
	"encoding/json"
	"github.com/DudeWhoCode/kulay/backend"
	"testing"
	"time"
)

func TestPut(t *testing.T) {
	host := "localhost"
	port := "6379"
	pass := ""
	db := 0
	queue := "testput"
	testCnt := 5
	client := backend.NewRedisSession(host, port, pass, db)
	client.Del(queue)
	pipe := make(chan string, testCnt)
	type test struct {
		Name  string `json:"name"`
		Desc  string `json:"desc"`
		Url   string `json:"url"`
		Stars int    `json:"stars"`
	}
	testData := &test{
		"kulay",
		"High speed message routing between services",
		"https://github.com/kulay",
		135,
	}
	testStr, _ := json.Marshal(testData)
	for i := 1; i <= testCnt; i++ {
		pipe <- string(testStr)
	}
	go Put(host, port, pass, db, queue, pipe)
	// Wait until Put populates redis queue
	// TODO : Implement timeout for channels
	time.Sleep(1 * time.Second)
	queueLen, err := client.LLen(queue).Result()
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if int(queueLen) != testCnt {
		t.Errorf("Expected message count in redis queue is %v, got %v", testCnt, queueLen)
	}

}

func TestGet(t *testing.T) {
	host := "localhost"
	port := "6379"
	pass := ""
	db := 0
	queue := "testget"
	testCnt := 5
	client := backend.NewRedisSession(host, port, pass, db)
	client.Del(queue)
	pipe := make(chan string, testCnt)
	type test struct {
		Name  string `json:"name"`
		Desc  string `json:"desc"`
		Url   string `json:"url"`
		Stars int    `json:"stars"`
	}
	testData := &test{
		"kulay",
		"High speed message routing between services",
		"https://github.com/kulay",
		135,
	}
	testStr, _ := json.Marshal(testData)
	for i := 1; i <= testCnt; i++ {
		if err := client.RPush(queue, testStr).Err(); err != nil {
			t.Fatalf("Expected no error while pushing messages, got %s", err)
		}
	}
	Get(host, port, pass, db, queue, pipe)
	if len(pipe) != testCnt {
		t.Errorf("Expected message count in channel is %v, got %v", testCnt, len(pipe))
	}

}
