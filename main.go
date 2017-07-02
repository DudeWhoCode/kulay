package main

import (
	"github.com/DudeWhoCode/kulay/cmd"
	. "github.com/DudeWhoCode/kulay/logger"
)

func main() {
	Log.Info("Starting kulay app")
	cmd.Execute()
}
