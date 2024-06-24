package server

import (
	"blockchainSim/internal/db"
	"blockchainSim/internal/services/account"
	"log"
	"net/http"
)

type Server struct {
	accountHandler *account.Handler
}

func New() *Server {
	database := db.InitDb()
	accountStore := account.NewStore(database)
	return &Server{
		accountHandler: account.NewHandler(accountStore),
	}
}

func (s *Server) Run() {
	r := s.initRouter()

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
