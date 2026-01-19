package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
}

var configuarations Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env variables")
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("version is required")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service Name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http Port is required")
		os.Exit(1)
	}
	port, err := strconv.ParseInt(httpPort, 10, 64)

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("JWT secret key is required")
		os.Exit(1)
	}

	configuarations = Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: jwtSecretKey,
	}
}

func GetConfig() Config {
	LoadConfig()
	return configuarations
}
