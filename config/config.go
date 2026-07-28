package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}


var configurations Config 


func loadConfig() {
   err :=  godotenv.Load()

   if err != nil{
	fmt.Println("faild to load the env file",err)
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

	if  err  != nil{
		fmt.Println("port must be number")
				os.Exit(1)
	}

	configurations = Config{
		Version:     verson,
		ServiceName: serviceName,
		HttpPort:    int(portInt),
	}

}


func GetConfig() Config{
	loadConfig()
	return configurations
}