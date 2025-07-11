// App project main.go
package main

import (
	"log"

	"github.com/venuyeredla/pan-services/configs"
	"github.com/venuyeredla/pan-services/internal"
	"github.com/venuyeredla/pan-services/internal/config"
)

func main() {

	ctx, waitGroup, waitFunc := internal.ShutdownGracefully()
	logger := config.InitializeLogger()

	logger.Println("Started using shutdown gracefully")
	rawYaml := configs.GetConfigYaml()
	config, configErr := config.NewConfig(rawYaml) // create configuration from yaml file.
	if configErr != nil {
		log.Fatalf("Error loading configuration: %v", configErr)
	}
	router := internal.ApplicationConfig()
	httpServer := internal.NewHttpServer(config, router.Handler(), waitGroup) // Initialize HTTP server with configuration and wait group

	//host, _ := os.Hostname()
	//logger.Printf("Host %s, Page/Block size =%v \n", host, os.Getpagesize())

	httpServer.StartUp(ctx.Done()) // Start the HTTP server with graceful shutdown

	waitFunc() // Wait for the shutdown signal
}
