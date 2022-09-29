package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler interface {
	Home()
	AllClients()
	ClientById()
}

func (s *Server) Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The server is ready to serve"))
	}
}

func (s *Server) AllClients() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cash := s.database.Cashe()
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}
		for _, value := range cash {
			clients, err := json.Marshal(value)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintf(w, "%s\n", clients)
		}
	}
}

func (s *Server) ClientById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cash := s.database.Cashe()
		info := mux.Vars(r)
		for key, value := range cash {
			if key == info["id"] {
				client, err := json.Marshal(value)
				if err != nil {
					log.Fatal(err)

				}
				fmt.Fprintf(w, "%s", client)
				return
			}
		}
		w.Write([]byte("There is no client with this id"))
	}
}
