package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// ConnectMongoDB estabelece a conexão com o MongoDB
func ConnectMongoDB() {
	uri := "mongodb://balta:e296cd9f@localhost:27017/"

	log.Println("Tentando conectar ao MongoDB...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Erro ao conectar ao MongoDB:", err)
	}

	// Verifica a conexão
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Erro ao pingar no MongoDB:", err)
	}

	log.Println("Conexão estabelecida com sucesso!")
	Client = client
}

// GetCollection retorna uma coleção do banco de dados MongoDB
func GetCollection(database, collection string) *mongo.Collection {
	if Client == nil {
		log.Fatal("MongoDB client is not initialized 2")
	}
	return Client.Database(database).Collection(collection)
}
