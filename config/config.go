package config

import (
	"github.com/spf13/viper"
	. "naren/kulay/logger"
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
	viper.SetDefault("sqs.delete", true)
}

// Parse kulay config
func Parse(service string, section string) ( interface{}, error) {
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
	}
	return nil, err
}

// Load configuration
func Load(Flag string) (cfg interface{}) {
	viperCfg()
	svc := strings.Split(Flag, ".")[0]
	sec := strings.Split(Flag, ".")[1]
	cfg, err := Parse(svc, sec)
	Log.Println("config.go cfg value : ", cfg)
	if err != nil {
		panic(err)
	}
	return
}
