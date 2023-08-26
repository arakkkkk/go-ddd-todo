package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"todo/config"
	"todo/ent"
	"todo/internal/server/di"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	cfg    *config.Config
	ent    *ent.Client
	router *chi.Mux
}

type Options func(opts *Server) error

func New() *Server {
	cfg := config.New()
	return &Server{
		cfg:    cfg,
		ent:    di.NewEnt(cfg),
		router: chi.NewRouter(),
	}
}

func (s *Server) Init() {
	s.newEnt()
	s.InitDomains()
}

func (s *Server) Run() {
	http.ListenAndServe(s.cfg.Api.Port, s.router)
}

func (s *Server) newEnt() *ent.Client {
	dsn := fmt.Sprintf("%s://%s:%s/%s?sslmode=%s&user=%s&password=%s",
		s.cfg.Database.Driver,
		s.cfg.Database.Host,
		s.cfg.Database.Port,
		s.cfg.Database.Name,
		s.cfg.Database.User,
	)
	client, err := ent.Open(s.cfg.Database.Driver, dsn)
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err != nil {
		log.Fatalf("failed opening connection: %v", err)
		panic(err)
	}
	return client
}
