package config

import (
	"os"
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	os.Setenv("KULAY_CONF", "../testdata/kulay.toml")
	flag := "sqs.us_queue"
	cfg := Load(flag)
	sqsCfg, ok := cfg.(SQSConf)
	if !ok {
		t.Errorf("expected type: %s, got: %s", "config.SQSconf", reflect.TypeOf(sqsCfg).String())
	}
}

func TestParseSQS(t *testing.T) {
	os.Setenv("KULAY_CONF", "../testdata/kulay.toml")
	svc := "sqs"
	sec := "us_queue"
	expectedUrl := "https://sqs.us-east-1.amazonaws.com/123456789/test"
	expectedRegion := "us-east-1"
	expectedDel := true
	viperCfg()
	cfg, err := Parse(svc, sec)
	if err != nil {
		panic(err)
	}
	sqsCfg := cfg.(SQSConf)
	receivedURL := sqsCfg.QueueUrl
	receivedRegion := sqsCfg.Region
	receivedDel := sqsCfg.Delete
	if receivedURL != expectedUrl {
		t.Errorf("expected queue URL: %s, got: %s", expectedUrl, receivedURL)
	} else if receivedRegion != expectedRegion {
		t.Errorf("expected region: %s, got: %s", expectedRegion, receivedRegion)
	} else if receivedDel != expectedDel {
		t.Errorf("expected delete flag: %t, got: %t", expectedDel, receivedDel)
	}

	svc = "sqs"
	sec = "sg_queue"
	expectedUrl = "https://sqs.ap-southeast-1.amazonaws.com/123456789/test"
	expectedRegion = "ap-southeast-1"
	expectedDel = false
	viperCfg()
	cfg, err = Parse(svc, sec)
	if err != nil {
		panic(err)
	}
	sqsCfg = cfg.(SQSConf)
	receivedURL = sqsCfg.QueueUrl
	receivedRegion = sqsCfg.Region
	receivedDel = sqsCfg.Delete
	if receivedURL != expectedUrl {
		t.Errorf("expected queue URL: %s, got: %s", expectedUrl, receivedURL)
	} else if receivedRegion != expectedRegion {
		t.Errorf("expected region: %s, got: %s", expectedRegion, receivedRegion)
	} else if receivedDel != expectedDel {
		t.Errorf("expected delete flag: %t, got: %t", expectedDel, receivedDel)
	}

}

func TestParseRedisQ(t *testing.T) {
	os.Setenv("KULAY_CONF", "../testdata/kulay.toml")
	svc := "redisq"
	sec := "localbuffer"
	expectedHost := "localhost"
	expectedPort := "6379"
	expectedPass := ""
	expectedDB := 0
	expectedQ := "test"
	viperCfg()
	cfg, err := Parse(svc, sec)
	if err != nil {
		panic(err)
	}
	redisCfg := cfg.(RedisqConf)
	receivedHost := redisCfg.Host
	receivedPort := redisCfg.Port
	receivedPass := redisCfg.Pass
	receivedDB := redisCfg.DB
	receivedQ := redisCfg.Queue

	if receivedHost != expectedHost {
		t.Errorf("expected host: %s, got: %s", expectedHost, receivedHost)
	} else if receivedPort != expectedPort {
		t.Errorf("expected port: %s, got: %s", expectedPort, receivedPort)
	} else if receivedPass != expectedPass {
		t.Errorf("expected password: %t, got: %t", expectedPass, receivedPass)
	} else if receivedDB != expectedDB {
		t.Errorf("expected DB: %t, got: %t", expectedDB, receivedDB)
	} else if receivedQ != expectedQ {
		t.Errorf("expected queue: %t, got: %t", expectedQ, receivedQ)
	}
}

func TestParseRedisPubSub(t *testing.T) {
	os.Setenv("KULAY_CONF", "../testdata/kulay.toml")
	svc := "redispubsub"
	sec := "localchannel"
	expectedHost := "localhost"
	expectedPort := "6379"
	expectedPass := ""
	expectedDB := 0
	expectedChan := "test_pubsub"
	viperCfg()
	cfg, err := Parse(svc, sec)
	if err != nil {
		panic(err)
	}
	redisCfg := cfg.(RedisPubsubConf)
	receivedHost := redisCfg.Host
	receivedPort := redisCfg.Port
	receivedPass := redisCfg.Pass
	receivedDB := redisCfg.DB
	receivedChan := redisCfg.Channel

	if receivedHost != expectedHost {
		t.Errorf("expected host: %s, got: %s", expectedHost, receivedHost)
	} else if receivedPort != expectedPort {
		t.Errorf("expected port: %s, got: %s", expectedPort, receivedPort)
	} else if receivedPass != expectedPass {
		t.Errorf("expected password: %t, got: %t", expectedPass, receivedPass)
	} else if receivedDB != expectedDB {
		t.Errorf("expected DB: %t, got: %t", expectedDB, receivedDB)
	} else if receivedChan != expectedChan {
		t.Errorf("expected channel: %t, got: %t", expectedChan, receivedChan)
	}
}

func TestParseRedisJsonl(t *testing.T) {
	os.Setenv("KULAY_CONF", "../testdata/kulay.toml")
	svc := "jsonl"
	sec := "local_backup"
	expectedPath := "/tmp/log_backup.jsonl"
	expectedRotate := true
	expectedBatch := 100
	viperCfg()
	cfg, err := Parse(svc, sec)
	if err != nil {
		panic(err)
	}
	jsonlCfg := cfg.(JsonlConf)
	receivedPath := jsonlCfg.Path
	receivedRotate := jsonlCfg.Rotate
	receivedBatch := jsonlCfg.Batch

	if receivedPath != expectedPath {
		t.Errorf("expected path: %s, got: %s", expectedPath, receivedPath)
	} else if receivedRotate != expectedRotate {
		t.Errorf("expected rotate flag: %s, got: %s", expectedRotate, receivedRotate)
	} else if receivedBatch != expectedBatch {
		t.Errorf("expected batch count: %t, got: %t", expectedBatch, receivedBatch)
	}
}
