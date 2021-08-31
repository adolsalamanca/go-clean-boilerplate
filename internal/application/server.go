package application

import (
	"github.com/adolsalamanca/go-clean-boilerplate/internal/domain"
	_interface "github.com/adolsalamanca/go-clean-boilerplate/internal/interface"
	"github.com/gorilla/mux"
)

type Server struct {
	router    *mux.Router
	service   Servicer
	logger    _interface.Logger
	collector _interface.MetricsCollector
	// tracing
}

type Servicer interface {
	GetItems() ([]domain.Item, error)
	CreateItem(i domain.Item) error
}

func NewServer(service Servicer, logger _interface.Logger, collector _interface.MetricsCollector) *Server {
	return &Server{
		collector: collector,
		logger:    logger,
		router:    mux.NewRouter(),
		service:   service,
	}
}
