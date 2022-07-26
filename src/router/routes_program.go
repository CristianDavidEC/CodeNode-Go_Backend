package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routing() *chi.Mux {
	routerMux := chi.NewMux()

	routerMux.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	routerMux.Post("/saveProgram", nil)
	routerMux.Get("/getAllPrograms", nil)
	routerMux.Get("/getProgram", nil)
	routerMux.Get("/runProgram", nil)

	return routerMux
}
