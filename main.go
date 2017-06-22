package main

import (
	"naren/kulay/cmd"
	"naren/kulay/config"
)

func main() {
	cmd.Execute()
	config.Load()
	//if err := cmd.RootCmd.Execute(); err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
}
