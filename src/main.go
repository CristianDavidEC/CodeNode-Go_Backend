package main

import (
	"codenode/packages/src/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	DATABASE_CONNECTION := os.Getenv("DATABASE_CONNECTION")
	API_KEY := os.Getenv("API_KEY")

	routerMux := router.Routing()
	server := NewServer(routerMux)
	server.Run(&Config{
		databaseConecction: DATABASE_CONNECTION,
		apiKey:             API_KEY,
	})
}
