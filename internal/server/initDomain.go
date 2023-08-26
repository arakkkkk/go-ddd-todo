package server

import (
	todoHandler "todo/internal/domain/todo/handler"
	todoRepository "todo/internal/domain/todo/repository"
	todoUsecase "todo/internal/domain/todo/usecase"
)

func (s *Server) InitDomains() {
	s.initTodo()
}

func (s *Server) initTodo() {
	todoRepository := todoRepository.New(s.ent)

	todoUseCase := todoUsecase.New(
		todoRepository,
	)

	todoHandler.RegisterHTTPEndPoints(s.router, todoUseCase)
}
