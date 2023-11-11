package dbase

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Photo     string             `json:"photo,omitempty" bson:"photo,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	Fname     string             `json:"fname,omitempty" bson:"fname,omitempty"`
	Lname     string             `json:"lname,omitempty" bson:"lname,omitempty"`
}

type Note struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	OwnerID   primitive.ObjectID `json:"owner_id" bson:"owner_id"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Content   string             `json:"content,omitempty" bson:"content,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Photos    []string           `json:"photos,omitempty" bson:"photos,omitempty"`
}

type Meta struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	SiteName string             `json:"title,omitempty" bson:"title,omitempty"`
}
