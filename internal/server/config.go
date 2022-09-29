package server

import "github.com/VladislavFirsov/Publisher-subscriber/internal/database"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Database *database.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: "8000",
		LogLevel: "debug",
		Database: database.NewConfig(),
	}
}
