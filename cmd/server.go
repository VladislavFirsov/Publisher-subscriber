package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/VladislavFirsov/Publisher-subscriber/internal/server"
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
	if err := server.Start(); err != nil {
		log.Fatalln(err)
	}

}
