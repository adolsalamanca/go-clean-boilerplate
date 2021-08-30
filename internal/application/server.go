package application

import (
	"github.com/adolsalamanca/go-clean-boilerplate/internal/domain/entities"
	"github.com/gorilla/mux"
)

type FieldType uint8

type Field struct {
	Key    string
	String string
	Int    int
}

func NewFieldString(key, value string) Field {
	return Field{
		Key:    key,
		String: value,
	}
}

func NewFieldInt(key string, value int) Field {
	return Field{
		Key: key,
		Int: value,
	}
}

type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
}

type Servicer interface {
	GetItems() ([]entities.Item, error)
	CreateItem(i entities.Item) error
}

type Server struct {
	router  *mux.Router
	service Servicer
	logger  Logger
	// collector
	// tracing
}

func NewServer(service Servicer, logger Logger) *Server {
	return &Server{
		logger:  logger,
		router:  mux.NewRouter(),
		service: service,
	}
}
