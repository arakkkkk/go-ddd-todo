package config

import (
	"os"
)

type Database struct {
	Driver                 string
	Host                   string
	Port                   string
	Name                   string
	User                   string
}

func DataStore() *Database {
	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")

	db := &Database{
		Driver: driver,
		Host: host,
		Port: port,
		Name: name,
		User: user,

	}
	return db
}

