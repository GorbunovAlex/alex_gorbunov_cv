package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Image       string             `json:"image" bson:"image"`
	Techstack   []string           `json:"techstack" bson:"techstack"`
}
