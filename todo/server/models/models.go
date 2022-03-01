package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	TodoId   int64              `json:"id,omitempty"`
	Todo     string             `json:"todo,omitempty"`
	IsDone   bool               `json:"isDone,omitempty"`
	Deadline time.Time          `json:"deadline,omitempty"`
}
