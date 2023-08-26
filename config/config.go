package config

type Config struct {
	Api
	Database
}

func New() *Config {
	return &Config{
		Api:           *newAPI(),
		Database:      *newDatabase(),
	}
}

