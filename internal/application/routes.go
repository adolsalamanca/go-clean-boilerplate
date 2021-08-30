package application

import (
	"encoding/json"
	"net/http"

	"github.com/adolsalamanca/go-clean-boilerplate/internal/domain/entities"
	"github.com/gorilla/mux"
)

func (s *Server) Routes() *mux.Router {
	s.router.HandleFunc("/health", s.Health()).Methods(http.MethodGet)
	s.router.HandleFunc("/items", s.GetItems()).Methods(http.MethodGet)
	s.router.HandleFunc("/items", s.CreateItem()).Methods(http.MethodPost)

	return s.router
}

func (s *Server) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Health is working"))
	}
}

func (s *Server) GetItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		i, err := s.service.GetItems()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		out, err := json.Marshal(i)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(out)

	}
}

func (s *Server) CreateItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		i := entities.Item{}
		d := json.NewDecoder(r.Body)
		err := d.Decode(&i)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad request"))
			return
		}

		err = s.service.CreateItem(i)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something wrong happened creating"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("All good creating"))

	}
}
