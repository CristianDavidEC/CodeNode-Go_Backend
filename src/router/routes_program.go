package router

import (
	"codenode/packages/src/api"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Routing() *chi.Mux {
	routerMux := chi.NewMux()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	routerMux.Use(
		middleware.Logger,
		middleware.Recoverer,
		cors.Handler,
	)

	routerMux.Get("/get-all-programs", api.GetAllPrograms)
	routerMux.Get("/get-program/{id}", api.GetProgram)
	routerMux.Post("/run-code", api.RunProgram)
	routerMux.Post("/save-program", api.SaveProgram)

	return routerMux
}
