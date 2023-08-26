package handler

import (
	"encoding/json"
	"net/http"
	"time"
	"todo/internal/domain/todo"
	"todo/internal/domain/todo/usecase"
	"todo/internal/util/response"
)

type Handler struct {
	usecase *usecase.TodoUsecase
}

func NewHandler(usecase *usecase.TodoUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	todos, err := h.usecase.List(r.Context())
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Bad request.")
	}
	response, err := json.MarshalIndent(todos, "", "")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

type CreateRequest struct {
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	Priority  int       `json:"priority"`
	CreatedAt time.Time `json:"create_at"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Bad request.")
		return
	}
	todo, err := h.usecase.Create(r.Context(), &todo.Schema{
		Title: req.Title,
		Completed: req.Completed,
		Priority: req.Priority,
		CretedAt: req.CreatedAt,
	})
	resp, err := json.MarshalIndent(todo, "", "")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Bad request.")
		return
	}
	w.Write([]byte(resp))
}
