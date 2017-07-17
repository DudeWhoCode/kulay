package config

import (
	. "github.com/DudeWhoCode/kulay/logger"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type SQSConf struct {
	QueueUrl string
	Region   string
	Delete   bool
	Service  string
}

type JsonlConf struct {
	Path string
}

type RedisqConf struct {
	Host  string
	Port  string
	Pass  string
	DB    int
	Queue string
}

type RedisPubsubConf struct {
	Host    string
	Port    string
	Pass    string
	DB      int
	Channel string
}

func viperCfg() {
	filePath := os.Getenv("KULAY_CONF")
	dir, file := path.Split(filePath)
	file = strings.TrimSuffix(file, filepath.Ext(file))
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvPrefix("KULAY")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(replacer)
	viper.AddConfigPath(dir)
	viper.SetConfigName(file)
	viper.SetConfigType("toml")
}

// Parse kulay config
func Parse(service string, section string) (interface{}, error) {
	err := viper.ReadInConfig()
	if err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			Log.Warn("Running without config file")
		default:
			return nil, err
		}
	}
	switch service {
	case "sqs":
		SQSCfg := SQSConf{}
		subtree := "sqs." + section
		subv := viper.Sub(subtree)
		SQSCfg.QueueUrl = subv.GetString("queue_url")
		SQSCfg.Region = subv.GetString("region")
		SQSCfg.Delete = subv.GetBool("delete_msg")
		return SQSCfg, err
	case "jsonl":
		JsonlCfg := JsonlConf{}
		subtree := "jsonl." + section
		subv := viper.Sub(subtree)
		JsonlCfg.Path = subv.GetString("path")
		return JsonlCfg, err
	case "redisq":
		RedisqCfg := RedisqConf{}
		subtree := "redisq." + section
		subv := viper.Sub(subtree)
		RedisqCfg.Host = subv.GetString("host")
		RedisqCfg.Port = subv.GetString("port")
		RedisqCfg.Pass = subv.GetString("password")
		RedisqCfg.DB = subv.GetInt("database")
		RedisqCfg.Queue = subv.GetString("queue")
		return RedisqCfg, err
	case "redispubsub":
		RedisPubSubCfg := RedisPubsubConf{}
		subtree := "redispubsub." + section
		subv := viper.Sub(subtree)
		RedisPubSubCfg.Host = subv.GetString("host")
		RedisPubSubCfg.Port = subv.GetString("port")
		RedisPubSubCfg.Pass = subv.GetString("password")
		RedisPubSubCfg.DB = subv.GetInt("database")
		RedisPubSubCfg.Channel = subv.GetString("channel")
		return RedisPubSubCfg, err
	}
	return nil, err
}

// Load configuration
func Load(Flag string) (cfg interface{}) {
	viperCfg()
	svc := strings.Split(Flag, ".")[0]
	sec := strings.Split(Flag, ".")[1]
	cfg, err := Parse(svc, sec)
	if err != nil {
		panic(err)
	}
	return
}
