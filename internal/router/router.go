package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/sjc5/go-api-template/internal/handler"
	"github.com/sjc5/go-api-template/internal/middleware"
)

func Init() *chi.Mux {
	r := chi.NewRouter()
	middleware.ApplyGlobal(r)
	setupPublicRoutes(r)
	setupProtectedRoutes(r)
	return r
}

func setupPublicRoutes(r chi.Router) {
	r.Get("/public", handler.Public)
}

func setupProtectedRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		applyProtectedMiddleware(r)
		r.Get("/protected", handler.Protected)
	})
}

func applyProtectedMiddleware(r chi.Router) {
	r.Use(middleware.RequireSession)
}
