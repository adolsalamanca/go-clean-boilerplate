package application

import (
	_interface "github.com/adolsalamanca/go-rest-boilerplate/internal/interface"
	"github.com/gorilla/mux"
)

type Server struct {
	router  *mux.Router
	service *_interface.Service
}

func NewServer(svc *_interface.Service) *Server {
	return &Server{
		router:  mux.NewRouter(),
		service: svc,
	}
}
