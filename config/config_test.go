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

func TestParse(t *testing.T) {
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
