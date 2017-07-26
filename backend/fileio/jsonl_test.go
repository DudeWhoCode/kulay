package fileio

import (
	"bufio"
	"encoding/json"
	"os"
	"testing"
	"path/filepath"
	"strings"
)

func TestPut(t *testing.T) {
	path := "/tmp/jsonltest/"
	os.RemoveAll(path)
	os.MkdirAll(path, 0777)
	fpath := path + "test_put.jsonl"
	pipe := make(chan string)
	batch := 5
	rotate := false
	testCnt := 10
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
	go Put(fpath, batch, pipe, rotate)
	for i := 1; i <= testCnt; i++ {
		pipe <- string(testStr)
	}
	fileList := []string{}
	filepath.Walk(path, func (path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})
	fpath = fileList[1]
	file, err := os.Open(fpath)
	if err != nil {
		t.Errorf("Expected no errors in reading file, got %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := json.Unmarshal(scanner.Bytes(), &test{}); err != nil {
			t.Errorf("Expected no errors in unmarshalling jsonline, got %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		t.Errorf("Expected no errors in scanning file, got %v", err)
	}
}

func TestGet(t *testing.T) {
	fpath := "/tmp/test_get.jsonl"
	testCnt := 10
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
	testMsg, _ := json.Marshal(testData)
	testMsg = append(testMsg, "\n"...)
	toWrite, err := os.Create(fpath)
	if err != nil {
		t.Fatal("Unable to open file for writing jsonl")
	}
	for i := 1; i <= testCnt; i++ {
		toWrite.Write(testMsg)
	}
	toWrite.Close()
	Get(fpath, pipe)
	if len(pipe) != testCnt {
		t.Errorf("Expected message count is %v, got %v", testCnt, len(pipe))
	}
}
