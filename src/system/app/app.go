package app

import (
	"log"
	"net/http"
)

type Server struct {
	port string
}

func NewServer() Server {
	return Server{}
}

func (s *Server) Init() {
	log.Println("Initializing server....")
	s.port = ":8000"

}

func (s *Server) Start() {
	log.Println("Starting server....")
	http.ListenAndServe(s.port, nil)
}
