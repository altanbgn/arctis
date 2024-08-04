package account

import (
  "time"
)

type CreateAccountPayload struct {
	Firstname string    `bson:"firstname" json:"firstname" validate:"required"`
	Lastname  string    `bson:"lastname" json:"lastname" validate:"required"`
	Email     string    `bson:"email" json:"email" validate:"required,email"`
	Username  string    `bson:"username" json:"username" validate:"required"`
	Password  string    `bson:"password" json:"password" validate:"required"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type UpdateAccountPayload struct {
	Firstname string    `bson:"firstname" json:"firstname" validate:"required"`
	Lastname  string    `bson:"lastname" json:"lastname" validate:"required"`
	Username  string    `bson:"username" json:"username" validate:"required"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
