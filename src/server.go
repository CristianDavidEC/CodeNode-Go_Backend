package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type ServerProgram struct {
	server *http.Server
}

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

func (s *ServerProgram) Run() {
	fmt.Println("Running on port 3080")
	log.Fatal(s.server.ListenAndServe())
}
