package sqsapp

import (
	"testing"
	"encoding/json"
)

func TestGet(t *testing.T){
	testCnt := 50
	region := "ap-southeast-1"
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
	qURL := "https://sqs.ap-southeast-1.amazonaws.com/971037846030/test_2"
	for i := 0; i < testCnt; i++ {
		producerPipe <- string(testStr)
	}
	t.Logf("Sent %v messages to producer channel for testing", len(producerPipe))
	close(producerPipe)
	Put(qURL, region, producerPipe)
	consumerPipe := make(chan string, testCnt)
	Get(qURL, region, true, consumerPipe)
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
	if len(testResults) != testCnt{
		t.Errorf("Expected consumed message count is 50, got %v",testResults)
	}
	t.Logf("Unpacked %v messages from consumer channel", len(testResults))

}
