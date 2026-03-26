package config

import (
	"fmt"
	"io"
	"log"
	"os"
)

// create constant for log file path
const LogFilePath = "logs/log.log"

func InitializeLogger() *log.Logger {
	logFile, err := os.OpenFile(LogFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("\nUnable to create log file in the path=%s\n", LogFilePath)
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	logger := log.Default()
	return logger
}
