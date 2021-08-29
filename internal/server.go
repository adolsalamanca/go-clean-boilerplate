package internal

import (
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	svc    Service
}

func NewServer() *Server {
	return &Server{router: mux.NewRouter()}
}
