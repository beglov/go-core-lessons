package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

// API предоставляет интерфейс программного взаимодействия.
type API struct {
	router *mux.Router
}

// New создаёт объект API.
func New(r *mux.Router) *API {
	api := API{
		router: r,
	}
	return &api
}

// Endpoints регистрирует конечные точки API.
func (api *API) Endpoints() {
	api.router.Use(api.jwtMiddleware)
	api.router.HandleFunc("/api/v1/auth", api.authJWT).Methods(http.MethodPost, http.MethodOptions)
}
