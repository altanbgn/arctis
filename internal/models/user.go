package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Firstname string             `bson:"firstname,omitempty" json:"firstname,omitempty" validate:"required"`
	Lastname  string             `bson:"lastname,omitempty" json:"lastname,omitempty" validate:"required"`
	Bio       string             `bson:"bio,omitempty" json:"bio,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty" validate:"required,email"`
	Username  string             `bson:"username,omitempty" json:"username,omitempty" validate:"required"`
	Password  string             `bson:"password,omitempty" json:"password,omitempty" validate:"required"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
