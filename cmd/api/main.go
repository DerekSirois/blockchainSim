package main

import (
	"blockchainSim/internal/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	s := server.New()
	s.Run()
}
