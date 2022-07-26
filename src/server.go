package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type MyServer struct {
	server *http.Server
}

func NewServer(routerMux *chi.Mux) *MyServer {
	server := &http.Server{
		Addr:           ":3080",
		Handler:        routerMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &MyServer{server}
}

func (s *MyServer) Run() {
	fmt.Print("Running on port 3080")
	log.Fatal(s.server.ListenAndServe())
}
