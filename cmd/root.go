package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"naren/kulay/config"
	"strings"
	ksqs "naren/kulay/sqs"
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
	//RootCmd.AddCommand(FromCmd)
	RootCmd.PersistentFlags().StringVarP(&FromFlag, "from", "f", "",
		"Source service to route from")
	RootCmd.PersistentFlags().StringVarP(&ToFlag, "to", "t", "",
		"Source service to route to")
	//FromCmd.AddCommand(ToCmd)
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func kulayApp() {
	if FromFlag == "" || ToFlag == "" {
		fmt.Println("Need to specify both from and to flags")
		os.Exit(-1)
	}
	FromSvc := strings.Split(FromFlag, ".")[0]
	ToSvc := strings.Split(ToFlag, ".")[0]
	FromSec := strings.Split(FromFlag, ".")[1]
	ToSec := strings.Split(ToFlag, ".")[1]
	FromConfig := config.Load(FromSec)
	ToConfig := config.Load(ToSec)
	fmt.Println(FromSvc)
	fmt.Println(ToSvc)
	pipe := make(chan string, 20)
	done := make(chan bool)
	if FromSvc == "sqs" {
		ksqs.Consume(pipe, done, FromConfig)
	}
	if ToSvc == "sqs" {
		ksqs.Push(pipe, done, ToConfig)
	}
	<-done
}
