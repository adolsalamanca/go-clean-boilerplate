package application

import (
	"github.com/adolsalamanca/go-clean-boilerplate/internal/domain/entities"
	_interface "github.com/adolsalamanca/go-clean-boilerplate/internal/interface"
	"github.com/gorilla/mux"
)

type Server struct {
	router  *mux.Router
	service Servicer
	logger  _interface.Logger
	// collector
	// tracing
}

type Servicer interface {
	GetItems() ([]entities.Item, error)
	CreateItem(i entities.Item) error
}

func NewServer(service Servicer, logger _interface.Logger) *Server {
	return &Server{
		logger:  logger,
		router:  mux.NewRouter(),
		service: service,
	}
}
