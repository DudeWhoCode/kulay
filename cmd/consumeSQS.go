package cmd

import (
	"github.com/spf13/cobra"
	ksqs "naren/kulay/sqs"
	"fmt"
)

// consumeCmd represents the sqs consumer command
var ConsumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "sqs consumer",
	Long:  `sqs consumer`,
	Run: func(cmd *cobra.Command, args []string) {
		pipe := make(chan string, 20)
		done := make(chan bool)
		ksqs.Consume(pipe, done)
		ksqs.Push(pipe, done)
		<-done
		fmt.Println("DONE")
	},
}
