package cmd

import (
	"github.com/spf13/cobra"
	"naren/kulay/config"
	ksqs "naren/kulay/backend/sqs"
	"os"
	"strings"
	. "naren/kulay/logger"
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

func initFromSvc(svc string, cfg config.Kulay, pipe chan string, done chan bool) {
	switch svc {
	case "sqs":
		Log.Info("Initialized SQS consumer")
		ksqs.Get(pipe, done, cfg)
	}
}

func initToSvc(svc string, cfg config.Kulay, pipe chan string, done chan bool) {
	switch svc {
	case "sqs":
		Log.Info("Initialized SQS producer")
		ksqs.Put(pipe, done, cfg)
	}
}

func kulayApp() {
	if FromFlag == "" || ToFlag == "" {
		Log.Error("Need to specify both from and to flags")
		os.Exit(-1)
	}
	FromSvc := strings.Split(FromFlag, ".")[0]
	ToSvc := strings.Split(ToFlag, ".")[0]
	FromSec := strings.Split(FromFlag, ".")[1]
	ToSec := strings.Split(ToFlag, ".")[1]
	FromConfig := config.Load(FromSec)
	ToConfig := config.Load(ToSec)
	pipe := make(chan string, 20)
	done := make(chan bool)
	initFromSvc(FromSvc, FromConfig, pipe, done)
	initToSvc(ToSvc, ToConfig, pipe, done)
	<-done
}
