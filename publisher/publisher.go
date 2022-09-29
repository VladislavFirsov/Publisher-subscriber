package main

import (
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
)

func main() {
	file := []string{"model.json"}

	conn, err := stan.Connect("test-cluster", "publisher-id")
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	for _, file := range file {
		client, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		conn.Publish("clients", client)
	}
}
