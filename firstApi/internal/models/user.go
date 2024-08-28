package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User representa um usuário no sistema
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	CreatedAt time.Time          `bson: "createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson: "updatedAt" json:"updatedAt"`
}
