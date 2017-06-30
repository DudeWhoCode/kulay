package fileio

import (
	. "naren/kulay/logger"
	"os"
)

func Put(fpath string, rec <-chan string) {
	f, err := os.Create(fpath)
	if err != nil {
		Log.Error("Unable to open file for writing jsonl")
	}
	for msg := range rec {
		f.Write([]byte(msg + "\n"))
	}
	defer f.Close()
}
