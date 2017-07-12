package fileio

import (
	"bufio"
	. "github.com/DudeWhoCode/kulay/logger"
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

func Get(fpath string, snd chan<- string) {
	f, err := os.Open(fpath)
	Log.Info("file openend")
	if err != nil {
		Log.Error("Unable to open file for reading jsonl")
	}
	scanner := bufio.NewScanner(f)
	Log.Info("new scanner initiaited")
	for scanner.Scan() {
		snd <- string(scanner.Bytes())
		Log.Info("sending file content to channel")
	}
	if err := scanner.Err(); err != nil {
		Log.Fatal("Error while scanning the file\n", err)
	}
	defer f.Close()
}
