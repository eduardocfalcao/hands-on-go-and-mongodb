package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	versionHeader = "Accept"
	appName       = "sweatmgr"
)

// InitRouter - Initialize de routes
func InitRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	swaggerUIHandler := http.StripPrefix("/swaggerui", http.FileServer(http.Dir("./swaggerui/")))
	router.PathPrefix("/swaggerui").Handler(swaggerUIHandler)

	v1 := fmt.Sprintf("application/vnd.%s.v1", appName)

	router.HandleFunc("/sweat", createSweatHandler).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/sweat_samples", getSweatSamplesHandler).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/sweat/{id}", getSweatByIDHandler).Methods(http.MethodGet).Headers(versionHeader, v1)

	v2 := fmt.Sprintf("application/vnd.%s.v2", appName)

	router.HandleFunc("/users/{user_id}/sweat", getSweatByUserIDHandler).Methods(http.MethodGet).Headers(versionHeader, v2)

	return
}
