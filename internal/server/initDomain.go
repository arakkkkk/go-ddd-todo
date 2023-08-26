package server

import (
	todoRepository "todo/internal/domain/todo/repository"
	todoUsecase "todo/internal/domain/todo/usecase"
	todoHandler "todo/internal/domain/todo/handler"
)

func (s *Server) InitDomains() {
	s.initTodo()
}

func (s *Server) initTodo() {
	todoRepository := todoRepository.New(s.ent)

	newAuthorUseCase := todoUsecase.New(
		todoRepository,
	)

	todoHandler.RegisterHTTPEndPoints(s.router, newAuthorUseCase)
}
