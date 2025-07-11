package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/venuyeredla/pan-services/configs"
	"github.com/venuyeredla/pan-services/internal/config"
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	} else {
		fmt.Println("Environment variables loaded successfully")
	}
}

func main() {
	fmt.Println("Starting the application...", os.Getenv("SERVER_PORT"))
	rawYaml := configs.GetConfigYaml()
	config.NewConfig(rawYaml) // Load configuration from environment variable
}
