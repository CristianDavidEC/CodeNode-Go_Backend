package main

import (
	"codenode/packages/src/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	routerMux := router.Routing()
	server := NewServer(routerMux)
	server.Run()
}
