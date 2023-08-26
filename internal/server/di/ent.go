package di

import (
	"context"
	"fmt"
	"log"

	"todo/config"
	"todo/ent"

	_ "github.com/go-sql-driver/mysql"
)

func NewEnt(cfg *config.Config) *ent.Client {
	dsn := fmt.Sprintf("%s://%s:%s/%s?sslmode=%s&user=%s&password=%s",
		cfg.Database.Driver,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.User,
	)
	client, err := ent.Open(cfg.Database.Driver, dsn)
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err != nil {
		log.Fatalf("failed opening connection: %v", err)
		panic(err)
	}
	return client
}

