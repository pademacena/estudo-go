package routes

import (
	"firstApi/internal/handlers"

	"github.com/gorilla/mux"
)

// NewRouter cria um novo roteador com as rotas definidas
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Define a rota de "health check"
	router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	// Rota de registro de usuário
	UserRoutes(router)

	return router
}
