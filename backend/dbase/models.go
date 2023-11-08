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
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Email      string             `json:"email,omitempty" bson:"email,omitempty"`
	Photo      string             `json:"photo,omitempty" bson:"photo,omitempty"`
	Mobile     string             `json:"mobile,omitempty" bson:"mobile,omitempty"`
	Phone      string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Website    string             `json:"website,omitempty" bson:"website,omitempty"`
	X          string             `json:"x,omitempty" bson:"x,omitempty"`
	Instagram  string             `json:"instagram,omitempty" bson:"instagram,omitempty"`
	Facebook   string             `json:"facebook,omitempty" bson:"facebook,omitempty"`
	Address    string             `json:"address,omitempty" bson:"address,omitempty"`
	AvgPrice   float64            `json:"avgprice,omitempty" bson:"avgprice,omitempty"`
	NumOfEst   uint32             `json:"numofest,omitempty" bson:"numofest,omitempty"`
	Bio        string             `json:"bio,omitempty" bson:"bio,omitempty"`
	Rating     float64            `json:"rating,omitempty" bjson:"rating,omitempty" `
	AvgResTime float64            `json:"avgrestim,omitempty" bson:"avgrestim,omitempty"`
}

type Metadata struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Logo     string             `json:"logo,omitempty" bson:"logo,omitempty"`
	Carousal []string           `json:"carousal,omitempty" bson:"carousal,omitempty"`
}
