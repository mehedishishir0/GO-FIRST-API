package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config


type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
	
}


type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JWTSecret   string
	DB DBConfig
}




func loadConfig() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("faild to load the env file", err)
		os.Exit(1)
	}

	verson := os.Getenv("VERSION")

	if verson == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("service name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")

	if httpPort == "" {
		fmt.Println("http port name is required")
		os.Exit(1)
	}
	portInt, err := strconv.ParseInt(httpPort, 10, 64)

	if err != nil {
		fmt.Println("port must be number")
		os.Exit(1)
	}

	JWTSecret := os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		fmt.Println("JWT secret is required")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("DB host is required")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("DB port is required")
		os.Exit(1)
	}
	dbPortInt, err := strconv.ParseInt(dbPort, 10, 64)
	
	if err != nil {
		fmt.Println("DB port must be number")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("DB user is required")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("DB password is required")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB name is required")
		os.Exit(1)
	}

	dbSSLMode := os.Getenv("DB_SSL_MODE")
	if dbSSLMode == "" {
		fmt.Println("DB SSL mode is required")
		os.Exit(1)
	}

	configurations = &Config{
		Version:     verson,
		ServiceName: serviceName,
		HttpPort:    int(portInt),
		JWTSecret:   JWTSecret,
		DB: DBConfig{
			Host:     dbHost,
			Port:     int(dbPortInt),
			User:     dbUser,
			Password: dbPassword,
			Name:     dbName,
			SSLMode:  dbSSLMode,
		},
	}


}



func GetConfig() *Config {

	if configurations == nil {

		loadConfig()
	}
	return configurations
}
