package server

import (
	"log"
	"net/http"

	"firstApi/internal/db"
	"firstApi/internal/routes"
)

func Run() {
	// Conectar ao MongoDB
	log.Println("Iniciando a conexao com o Mongo")
	db.ConnectMongoDB("mongodb://balta:e296cd9f@localhost:27017/")

	// Configurar as rotas
	router := routes.NewRouter()

	// Inicia o servidor
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
