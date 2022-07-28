package router

import (
	"codenode/packages/src/api"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routing() *chi.Mux {
	routerMux := chi.NewMux()

	routerMux.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	routerMux.Post("/save-program", api.SaveProgram)
	routerMux.Get("/get-all-programs", api.GetAllPrograms)
	routerMux.Get("/get-program", api.GetProgram)
	routerMux.Post("/run-program", api.RunProgram)

	return routerMux
}
