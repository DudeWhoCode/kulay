package config

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"github.com/spf13/viper"
	"fmt"
)

type Kulay struct {
	QueueUrl string
	Region   string
	Delete   bool
}

// Kulay config variable
var KulayConf *Kulay

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

	//user, err := user.Current()
	//if err != nil {
	//	fmt.Println("{viperCfg}", err)
	//}

	//viper.SetDefault("queries.location", filepath.Join(user.HomeDir, "queries"))
}


// Parse kulay config
func Parse(cfg *Kulay) (err error) {
	err = viper.ReadInConfig()
	if err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			fmt.Println("Running without config file")
		default:
			return
		}
	}
	cfg.QueueUrl = viper.GetString("sqs.queue_url")
	cfg.Region = viper.GetString("sqs.region")
	cfg.Delete = viper.GetBool("sqs.delete_msg")
	return
}

// Load configuration
func Load() {
	viperCfg()
	KulayConf = &Kulay{}
	err := Parse(KulayConf)
	if err != nil {
		panic(err)
	}
	fmt.Println("LOADED CONFIG : ", KulayConf)
}