package application

import (
	"net/http"

	"github.com/adolsalamanca/go-rest-boilerplate/internal/domain/entities"
	"github.com/gorilla/mux"
)

func (s *Server) Routes() *mux.Router {

	s.router.Methods(http.MethodGet).Path("/health").HandlerFunc(s.Health())
	s.router.Methods(http.MethodGet).Path("/").HandlerFunc(s.GetItems())
	s.router.Methods(http.MethodPost).Path("/").HandlerFunc(s.CreateItem())

	return s.router
}

func (s *Server) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Health is good"))

	}
}

func (s *Server) GetItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, err := s.service.GetItems()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something wrong happened getting"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("All good getting"))

	}
}

func (s *Server) CreateItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := s.service.CreateItem(entities.Item{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something wrong happened creating"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("All good creating"))

	}

}
