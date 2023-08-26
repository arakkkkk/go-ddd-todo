package handler

import (
	"todo/internal/domain/todo/usecase"

	"github.com/go-chi/chi/v5"
)

func RegisterHTTPEndPoints(router *chi.Mux, usecase *usecase.TodoUsecase) {
	h := NewHandler(usecase)
	router.Route("/api/v1/todo", func(router chi.Router) {
		router.Post("/create", h.Create)
		router.Get("/list", h.List)
	})
}
