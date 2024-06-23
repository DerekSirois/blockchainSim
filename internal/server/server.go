package server

import (
	"blockchainSim/internal/db"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Server struct {
	db *sqlx.DB
}

func New() *Server {
	return &Server{
		db: db.InitDb(),
	}
}

func (s *Server) Run() {
	r := s.initRouter()

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
