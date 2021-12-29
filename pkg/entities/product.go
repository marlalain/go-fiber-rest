package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title" binding:"required"`
	Description string             `json:"description" bson:"description"`
}

type DeleteRequest struct {
	ID string `json:"id" binding:"required,gte=1"`
}
