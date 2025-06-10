// App project main.go
package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/venuyeredla/pan-services/internal"
	"github.com/venuyeredla/pan-services/internal/repository"
)

const (
	port = "2024"
)

func main() {
	logger := ConfigLogger()
	host, _ := os.Hostname()
	logger.Printf("Host %s, Page/Block size =%v \n", host, os.Getpagesize())
	router := internal.ApplicationConfig()
	//ginEngine.Run()
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router.Handler(),
	}
	go func() {
		// service connections
		logger.Printf("Application is available at Host: %v and  Port=%v", host, port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error : %s\n", err)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	logger.Println("Shutting down db connection pool")
	repository.ClosePool()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}

func ConfigLogger() *log.Logger {
	logFile, err := os.OpenFile("logs/log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	logger := log.Default()
	return logger
}
