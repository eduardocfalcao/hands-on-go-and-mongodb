package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

// InitRouter - Initialize de routes
func InitRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)
	return
}
