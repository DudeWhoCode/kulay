package sqsapp

import (
	"encoding/json"
	"testing"
)

func TestSQS(t *testing.T) {
	testCnt := 5
	region := "us-east-1"
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
	var testResults []*test
	testStr, _ := json.Marshal(testData)
	producerPipe := make(chan string, testCnt)
	qURL := "http://localhost:3000/123456789012/test"
	for i := 0; i < testCnt; i++ {
		producerPipe <- string(testStr)
	}
	t.Logf("Sent %v messages to producer channel for testing", len(producerPipe))
	close(producerPipe)
	Put(qURL, region, producerPipe, true)
	consumerPipe := make(chan string, testCnt)
	Get(qURL, region, true, consumerPipe, true)
	t.Logf("Received %v messages from SQS to consumer channel", len(consumerPipe))
	close(consumerPipe)
	for msg := range consumerPipe {
		testResult := &test{}
		if err := json.Unmarshal([]byte(msg), testResult); err != nil {
			t.Errorf("Expected no errors in unmarshalling jsonline, got %v", err)
		} else {
			testResults = append(testResults, testResult)
		}
	}
	if len(testResults) != testCnt {
		t.Errorf("Expected consumed message count is %v, got %v", testCnt, testResults)
	}
	t.Logf("Unpacked %v messages from consumer channel", len(testResults))

}

func TestRegions(t *testing.T) {
	testCnt := 5
	region := "us-east-1"
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
	var testResults []*test
	testStr, _ := json.Marshal(testData)
	producerPipe := make(chan string, testCnt)
	qURL := "http://localhost:3000/123456789012/test"
	destqURL := "http://localhost:3000/123456789012/test"
	destRegion := "ap-southeast-1"
	for i := 0; i < testCnt; i++ {
		producerPipe <- string(testStr)
	}
	t.Logf("Sent %v messages to producer channel for testing", len(producerPipe))
	close(producerPipe)
	Put(qURL, region, producerPipe, true)
	consumerPipe := make(chan string, testCnt)
	Get(qURL, region, true, consumerPipe, true)
	t.Logf("Received %v messages from SQS to consumer channel", len(consumerPipe))
	close(consumerPipe)
	Put(destqURL, destRegion, consumerPipe, true)
	resultPipe := make(chan string, testCnt)
	Get(destqURL, destRegion, true, resultPipe, true)
	if len(resultPipe) != testCnt {
		t.Errorf("Expected consumed message count is %v, got %v", testCnt, testResults)
	}
	t.Logf("Tested cross region sqs producer and consumer")

}
