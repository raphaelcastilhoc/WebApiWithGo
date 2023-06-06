package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string
	Quantity int
}
