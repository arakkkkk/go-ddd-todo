package config

import (
	"log"
)

type Config struct {
	Api
	Cors

	Database
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	return &Config{
		Api:           API(),
		Cors:          NewCors(),
		Database:      DataStore(),
		Cache:         NewCache(),
		Elasticsearch: ElasticSearch(),
		Session:       NewSession(),
	}
}

