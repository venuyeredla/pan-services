package config

import (
	"io"
	"log"
	"os"
)

func InitializeLogger() *log.Logger {
	logFile, err := os.OpenFile("logs/log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	logger := log.Default()
	return logger
}
