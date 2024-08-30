package routes

import (
	"firstApi/internal/handlers"

	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/register", handlers.RegisterUser).Methods("POST")

	return router
}
