package dbase

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Estate struct {
	ID        primitive.ObjectID     `json:"_id" bson:"_id"`
	TITLE     string                 `json:"title,omitempty" bson:"title,omitempty"`
	OwnerID   primitive.ObjectID     `json:"owner_id" bson:"owner_id"`
	Latitude  float64                `json:"latitude,omitempty" bson:"latitude,omitempty"`
	Longitude float64                `json:"longitude,omitempty" bson:"longitude,omitempty"`
	Price     float64                `json:"price,omitempty" bson:"price,omitempty"`
	Currency  string                 `json:"currency,omitempty" bson:"currency,omitempty"`
	Photos    []string               `json:"photos,omitempty" bson:"photos,omitempty"`
	Features  map[string]interface{} `json:"features,omitempty" bson:"features,omitempty"`
	CreatedAt time.Time              `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time              `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type Owner struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Email string             `json:"email,omitempty" bson:"email,omitempty"`
}
