package application

import (
	"github.com/adolsalamanca/go-clean-boilerplate/internal/domain/entities"
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

func NewServer(service Servicer) *Server {
	return &Server{
		router:  mux.NewRouter(),
		service: service,
	}
}
