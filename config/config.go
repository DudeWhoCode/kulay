package config

import (
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
	"strings"
	. "naren/kulay/logger"
)

type Kulay struct {
	QueueUrl string
	Region   string
	Delete   bool
	Service  string
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
func Parse(section string) (KulayConf Kulay, err error) {
	KulayConf = Kulay{}
	err = viper.ReadInConfig()
	if err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			Log.Warn("Running without config file")
		default:
			return
		}
	}
	subtree := "sqs." + section
	subv := viper.Sub(subtree)
	KulayConf.QueueUrl = subv.GetString("queue_url")
	KulayConf.Region = subv.GetString("region")
	KulayConf.Delete = subv.GetBool("delete_msg")
	return
}

// Load configuration
func Load(section string) Kulay {
	viperCfg()
	cfg, err := Parse(section)
	if err != nil {
		panic(err)
	}
	return cfg
}
