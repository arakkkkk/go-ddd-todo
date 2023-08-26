package config

import (
	"os"
)

type Api struct {
	Port              string
}

func newAPI() *Api {
	port := os.Getenv("API_PORT")
	api := &Api{
		Port: port,
	}

	return api
}

