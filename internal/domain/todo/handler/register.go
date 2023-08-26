package handler

import (
	"github.com/go-chi/chi/v5"
	"todo/internal/domain/todo/usecase"
)

func RegisterHTTPEndPoints(router *chi.Mux, usecase *usecase.TodoUsecase) {
	h := NewHandler(usecase)
	router.Route("/api/v1/todo", func(router chi.Router) {
		router.Post("/create", h.Create)
		router.Post("/List", h.List)
	})
}
