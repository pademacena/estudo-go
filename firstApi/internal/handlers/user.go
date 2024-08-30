package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"firstApi/internal/db"
	"firstApi/internal/models"
	"firstApi/internal/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func testConnection() {
	if db.Client != nil {
		userCollection = db.GetCollection("mydatabase", "users")
	} else {
		log.Fatal("MongoDB client is not initialized")
	}
}

// RegisterUser é o handler para registrar um novo usuário
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	testConnection()

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	if user.Name == "" || user.Email == "" || user.Password == "" {
		messageError := types.Error{
			Message: "Algum campo esta faltando",
		}
		log.Println("Erro")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(messageError)
		return
	}

	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if CheckIfEmailExists(ctx, user.Email) {
		messageError := types.Error{
			Message: "Email Ja registrado",
		}
		log.Println("Erro")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(messageError)
		return
	}

	_, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, "Erro ao registrar o usuário", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func CheckIfEmailExists(ctx context.Context, email string) bool {

	filter := bson.M{"email": email}
	var user models.User

	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false // E-mail não encontrado, pode inserir
		}
		return false // Outro erro ocorreu
	}

	return true // E-mail já existe
}
