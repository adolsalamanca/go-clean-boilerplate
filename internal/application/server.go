package application

import (
	"github.com/adolsalamanca/go-rest-boilerplate/internal/domain/entities"
	_interface "github.com/adolsalamanca/go-rest-boilerplate/internal/interface"
	"github.com/gorilla/mux"
)

type Servicer interface {
	GetItems() ([]entities.Item, error)
	CreateItem(i entities.Item) error
}

type Server struct {
	router  *mux.Router
	service Servicer
	// collector
	// tracing
	// logger
}

func NewServer(svc *_interface.Service) *Server {
	return &Server{
		router:  mux.NewRouter(),
		service: svc,
	}
}
