package server

import "github.com/gorilla/mux"

func (s *Server) initRouter() *mux.Router {
	r := mux.NewRouter()

	return r
}
