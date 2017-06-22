package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)
// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "kulay",
	Short: "High speed message routing",
	Long:  `Pull messages from desired service and push to other service or write to file system.
			SQS -> Redis, RedisPubSub -> SQS, RabbitMQ -> kafka ...`,
	Run: func(cmd *cobra.Command, args []string) {
		kulayApp()
	},
}

func Execute() {
	RootCmd.AddCommand(ConsumeCmd)
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func kulayApp()  {
	fmt.Println("Kulay App started")
}