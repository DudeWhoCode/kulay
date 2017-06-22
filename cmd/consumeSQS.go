package cmd

import (
	"github.com/spf13/cobra"
	ksqs "naren/kulay/sqs"
)

// consumeCmd represents the sqs consumer command
var ConsumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "sqs consumer",
	Long:  `sqs consumer`,
	Run: func(cmd *cobra.Command, args []string) {
		ksqs.Consume()
	},
}
