package main

import (
	"naren/kulay/cmd"
	. "naren/kulay/logger"
)

func main() {
	Log.Info("Starting kulay app")
	cmd.Execute()
}
