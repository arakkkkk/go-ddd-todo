package repository

import (
	"context"
	"fmt"
	"time"
	"todo/ent"
	"todo/internal/domain/todo"
)

type TodoRepository struct{
	ent *ent.Client
}

func New(ent *ent.Client) *TodoRepository {
	return &TodoRepository{
		ent: ent,
	}
}

func entBindSchema(t *ent.Todo) *todo.Schema {
	resp := &todo.Schema{
		Title: t.Title,
		Completed: t.Completed,
		Priority: t.Priority,
		CretedAt: t.CreatedAt,
	}
	return resp
}
func entTodosBindEntiryTodos(t []*ent.Todo) []*todo.Schema {
	resp := make([]*todo.Schema, 0)
	for _, v := range t {
		resp = append(resp, entBindSchema(v))
	}
	return resp
}

func (r *TodoRepository) List(ctx context.Context) ([]*todo.Schema, error) {
  todos, err := r.ent.Todo.Query().All(ctx)
  resp := entTodosBindEntiryTodos(todos)
  if err != nil {
		return nil, fmt.Errorf("failed query user: %v", err)
	}
  return resp, nil
}

func (r *TodoRepository) Create(ctx context.Context, t *todo.Schema) (*todo.Schema, error) {
  todo, err := r.ent.Todo.
    Create().
    SetTitle(t.Title).
    SetCompleted(t.Completed).
    SetPriority(t.Priority).
    SetCreatedAt(time.Now()).
    Save(ctx)

  if err != nil {
		return nil, fmt.Errorf("failed creating todo: %v", err)
	}

	entityTodo := entBindSchema(todo)
  return entityTodo, nil
}

