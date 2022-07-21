package main

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	setLevel()
	//setFormat()
	setOutput()
	//setMultiOutputs()

	levels()
}

func setLevel() {
	// level: only log the severity or above
	log.SetLevel(log.TraceLevel)
}

func setFormat() {
	// format: Text (default) OR JSON
	log.SetFormatter(&log.JSONFormatter{})
}

func setOutput() {
	// output: stderr (default) OR stdout, can be any io.Writer
	log.SetOutput(os.Stdout)
}

func setMultiOutputs() {
	writer1 := os.Stdout
	writer2, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}

	log.SetOutput(io.MultiWriter(writer1, writer2))
}

func levels() {
	log.Trace("This is trace")
	log.Debug("This is debug")
	log.Info("This is info") // default severity
	log.Warn("This is warn")
	log.Error("This is error")

	log.Fatal("This is fatal")
	//log.Panic("This is panic")
}
