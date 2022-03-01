package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	First_name    *string            `json:"firstname" validate:"required, min=2, max=100"`
	Last_name     *string            `json:"lastname" validate:"required, min=2, max=100"`
	Password      *string            `json:"password" validate:"required"`
	Email         *string            `json:"email" validate:"required"`
	Phone         *string            `json:"phone"`
	Token         *string            `json:"token"`
	User_type     *string            `json:"userType" validate:"required. eq=ADMIN|eq=USER"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    *time.Time         `json:"created_at"`
	Updated_at    *time.Time         `json:"updated_at"`
	User_id       string             `json:"user_id"`
}
