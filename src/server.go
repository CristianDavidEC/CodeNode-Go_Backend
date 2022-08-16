package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	database "codenode/packages/src/dataBase"

	"github.com/go-chi/chi/v5"
)

type ServerProgram struct {
	server *http.Server
}

type Config struct {
	databaseConecction string
	apiKey             string
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
func (s *ServerProgram) Run(config *Config) {
	clientDb := database.NewClient(config.databaseConecction, config.apiKey)
	database.SetTxnClient(clientDb)
	log.Fatal(s.server.ListenAndServe())
	fmt.Println("Running on port 3080")

}
