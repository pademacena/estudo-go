package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"firstApi/internal/db"
	"firstApi/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func init() {
	if db.Client != nil {
		userCollection = db.GetCollection("mydatabase", "users")
	} else {
		log.Fatal("MongoDB client is not initialized")
	}
}

// RegisterUser é o handler para registrar um novo usuário
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, "Erro ao registrar o usuário", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
