package database

import (
	"database/sql"
	"github.com/VladislavFirsov/Publisher-subscriber/internal/model"
	_ "github.com/lib/pq"
	"log"
)

var cash = make(map[string]model.Client)

type Database struct {
	config *Config
	pool   *sql.DB
	cash   map[string]model.Client
}

func NewDb(config *Config) *Database {
	return &Database{
		config: config,
		cash:   cash,
	}
}

func (d *Database) Open() error {
	pool, err := sql.Open("postgres", d.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := pool.Ping(); err != nil {
		return err
	}
	d.pool = pool
	return nil
}

func (d *Database) Close() {
	d.pool.Close()
}

func (d *Database) Cashe() map[string]model.Client {
	info, err := d.pool.Query("SELECT * from clients")
	if err != nil {
		log.Fatal(err)
	}
	for info.Next() {
		client := model.Client{}
		err := info.Scan(&client.ID, &client.Name, &client.Age, &client.City, &client.Phone, &client.Email)
		if err != nil {
			log.Fatal(err)
		}
		cash[client.ID] = client
	}
	return cash
}
