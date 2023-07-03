package main

import (
	"github.com/joho/godotenv"
	"github.com/qcodelabsllc/exag/email/network"
	"log"
)

func main() {
	// load config from env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// start gRPC server
	network.StartServer()
}
