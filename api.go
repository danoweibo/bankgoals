package main

import (
	"net/http"

	"github.com/gorilla/mux"
)
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
}

type apiFunc func (http.ResponseWriter, http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle the error
		}
	}
}

type APIServer struct {
	address string
}

func NewAPIServer(address string) *APIServer {
	return &APIServer{
		address: address,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", s.handleAccount())
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r http.Request) error {
	return nil
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r http.Request) error {
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r http.Request) error {
	return nil
}

func (s *APIServer) handleCloseAccount(w http.ResponseWriter, r http.Request) error {
	return nil
}

func (s *APIServer) handleTransferAccount(w http.ResponseWriter, r http.Request) error {
	return nil
}