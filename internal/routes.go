package internal

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) Routes() *mux.Router {

	s.router.Methods(http.MethodGet).HandlerFunc(s.Health())
	s.router.Methods(http.MethodPost).HandlerFunc(s.CreateItem())

	return s.router
}

func (s *Server) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("All good"))

	}
}

func (s *Server) GetItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("All good"))

	}
}

func (s *Server) CreateItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("All good"))

	}
}
