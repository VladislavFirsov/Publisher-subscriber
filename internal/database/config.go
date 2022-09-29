package database

type Config struct {
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		DatabaseURL: "user=postgres password=root dbname=postgres sslmode=disable",
	}
}
