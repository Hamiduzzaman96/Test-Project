package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

var configuarations Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env variables")
	}

	version := os.Getenv("VERSION")
	if version == "" {
		log.Fatal("version is required")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatal("Service Name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Fatal("Http Port is required")
		os.Exit(1)
	}
}

func GetConfig() Config {
	LoadConfig()
	return configuarations
}
