package main

import (
	"log"

	"github.com/deanonqq/microservice-user-balance/config"
	"github.com/deanonqq/microservice-user-balance/internal/apiserver"
	"github.com/joho/godotenv"
)

// init .env variables
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found!")
	}
}

func main() {
	conf := config.New()
	s := apiserver.New(conf)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
