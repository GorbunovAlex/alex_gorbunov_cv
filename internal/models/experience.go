package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Experience struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	Duration    string             `json:"duration" bson:"duration"`
	Description string             `json:"description" bson:"description"`
	Title       string             `json:"title" bson:"title"`
}
