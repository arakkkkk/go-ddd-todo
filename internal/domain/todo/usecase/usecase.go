package usecase

import (
	"context"
	"todo/internal/domain/todo"
	"todo/internal/domain/todo/repository"
)

type TodoUsecase struct {
	repo *repository.TodoRepository
}

func New(repo *repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		repo: repo,
	}
}

func (u *TodoUsecase) Create(ctx context.Context, r *todo.Schema) (*todo.Schema, error) {
	return u.repo.Create(ctx, r)
}

func (u *TodoUsecase) List(ctx context.Context) ([]*todo.Schema, error) {
	return u.repo.List(ctx)
}
