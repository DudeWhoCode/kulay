package fileio

import (
	"naren/kulay/config"
	"os"
	. "naren/kulay/logger"
)

func writeLine(fpath string, rec <-chan string, done chan bool) {
	f, err := os.Create(fpath)
	if err != nil {
		Log.Error("Unable to open file for writing jsonl")
	}
	for msg := range rec {
		f.Write([]byte(msg + "\n"))
	}
	defer f.Close()
	done <- true
}



func Put(pipe <-chan string, done chan bool, cfg interface{}) {
	jsonlCfg := cfg.(config.JsonlConf)
	qURL := jsonlCfg.Path
	writeLine(qURL, pipe, done)
}