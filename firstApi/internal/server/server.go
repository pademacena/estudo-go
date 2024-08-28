package server

import (
	"log"
	"net/http"

	"firstApi/internal/db"
	"firstApi/internal/routes"
)

func Run() {
	// Conectar ao MongoDB
	db.ConnectMongoDB()

	// Configurar as rotas
	router := routes.NewRouter()

	// Inicia o servidor
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
