package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	database "codenode/packages/src/dataBase"
	"codenode/packages/src/repository"

	"github.com/go-chi/chi/v5"
)

type ServerProgram struct {
	server *http.Server
}

/*Create the backend server*/
func NewServer(routerMux *chi.Mux) *ServerProgram {
	server := &http.Server{
		Addr:           ":3080",
		Handler:        routerMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &ServerProgram{server}
}

/*Execute the server*/
func (s *ServerProgram) Run() {
	DATABASE_CONNECTION := os.Getenv("DATABASE_CONNECTION")
	API_KEY := os.Getenv("API_KEY")
	clientDb, err := database.NewDgraphClient(DATABASE_CONNECTION, API_KEY)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(clientDb)
	fmt.Println("Running on port 3080")
	log.Fatal(s.server.ListenAndServe())
}
