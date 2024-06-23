package server

import (
	"log"
	"net/http"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (s *Server) Run() {
	r := s.initRouter()

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
