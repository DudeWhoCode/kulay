package main

import (
	"naren/kulay/cmd"
	"naren/kulay/config"
)

func main() {
	config.Load()
	cmd.Execute()
}
