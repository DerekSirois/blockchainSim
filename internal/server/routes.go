package server

import (
	"blockchainSim/internal/utils"

	"github.com/gorilla/mux"
)

func (s *Server) initRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/account", utils.Handler(s.accountHandler.Create)).Methods("POST")

	return r
}
