package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/VladislavFirsov/Publisher-subscriber/internal/model"
	"github.com/VladislavFirsov/Publisher-subscriber/internal/server"
	"github.com/nats-io/stan.go"
	"log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {
	flag.Parse()

	//creating config for server
	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatalln(err)
	}
	//creating new server

	server := server.New(config)
	var client model.Client
	conn, err := stan.Connect("test-cluster", "subscribe-id")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	sub, err := conn.Subscribe("clients", func(m *stan.Msg) {

		err := json.Unmarshal(m.Data, &client)
		if err != nil {
			log.Fatalln(err)
			return
		}
		pool, err := sql.Open("postgres", "user=postgres password=root dbname=postgres sslmode=disable")

		_, err = pool.Exec("INSERT INTO clients VALUES($1, $2, $3, $4, $5, $6)", client.ID, client.Name, client.Age, client.City, client.Phone, client.Email)
		if err != nil {
			log.Printf("DB execution error, %s", err.Error())
			return
		}

	})
	defer sub.Close()
	if err := server.Start(); err != nil {
		log.Fatalln(err)
	}

}
