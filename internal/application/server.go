package application

import (
	"github.com/adolsalamanca/go-rest-boilerplate/internal/domain/entities"
	_interface "github.com/adolsalamanca/go-rest-boilerplate/internal/interface"
	"github.com/gorilla/mux"
)

type Facader interface {
	GetItems() ([]entities.Item, error)
	CreateItem(i entities.Item) error
}

type Server struct {
	router *mux.Router
	facade Facader
	// collector
	// tracing
	// logger
}

func NewServer(facade *_interface.Facade) *Server {
	return &Server{
		router: mux.NewRouter(),
		facade: facade,
	}
}
