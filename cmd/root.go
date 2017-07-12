package cmd

import (
	jsonl "github.com/DudeWhoCode/kulay/backend/fileio"
	ksqs "github.com/DudeWhoCode/kulay/backend/sqs"
	redisq "github.com/DudeWhoCode/kulay/backend/redisq"
	"github.com/DudeWhoCode/kulay/config"
	. "github.com/DudeWhoCode/kulay/logger"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var FromFlag string
var ToFlag string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "kulay",
	Short: "High speed message routing",
	Long: `Pull messages from desired service and push to other service or write to file system.
			SQS -> Redis, RedisPubSub -> SQS, RabbitMQ -> kafka ...`,
	Run: func(cmd *cobra.Command, args []string) {
		kulayApp()
	},
}

func Execute() {
	RootCmd.PersistentFlags().StringVarP(&FromFlag, "from", "f", "",
		"Source service to route from")
	RootCmd.PersistentFlags().StringVarP(&ToFlag, "to", "t", "",
		"Source service to route to")
	if err := RootCmd.Execute(); err != nil {
		Log.Error("Command Execution error : ", err)
		os.Exit(-1)
	}
}

func initFromSvc(svc string, cfg interface{}, pipe chan string) {
	switch svc {
	case "sqs":
		Log.Info("Initialized SQS consumer")
		sqsCfg := cfg.(config.SQSConf)
		qURL := sqsCfg.QueueUrl
		del := sqsCfg.Delete
		region := sqsCfg.Region
		go ksqs.Get(qURL, region, del, pipe)
	case "jsonl":
		Log.Info("Initialized jsonl consumer")
		cfg := cfg.(config.JsonlConf)
		fPath := cfg.Path
		go jsonl.Get(fPath, pipe)
	}
}

func initToSvc(svc string, cfg interface{}, pipe chan string) {
	switch svc {
	case "sqs":
		Log.Info("Initialized SQS producer")
		sqsCfg := cfg.(config.SQSConf)
		qURL := sqsCfg.QueueUrl
		region := sqsCfg.Region
		go ksqs.Put(qURL, region, pipe)
	case "jsonl":
		Log.Info("Initialized jsonl producer")
		cfg := cfg.(config.JsonlConf)
		fPath := cfg.Path
		go jsonl.Put(fPath, pipe)
	case "redisq":
		Log.Info("Initialized redis producer")
		cfg := cfg.(config.RedisqConf)
		host := cfg.Host
		port := cfg.Port
		pass := cfg.Pass
		db := cfg.DB
		queue := cfg.Queue
		go redisq.Put(host, port, pass, db, queue, pipe)
	}
}

func kulayApp() {
	if FromFlag == "" || ToFlag == "" {
		Log.Error("Need to specify both from and to flags")
		os.Exit(-1)
	}
	FromSvc := strings.Split(FromFlag, ".")[0]
	ToSvc := strings.Split(ToFlag, ".")[0]
	FromConfig := config.Load(FromFlag)
	ToConfig := config.Load(ToFlag)
	pipe := make(chan string, 100)
	done := make(chan bool)
	initFromSvc(FromSvc, FromConfig, pipe)
	initToSvc(ToSvc, ToConfig, pipe)
	<-done
}
