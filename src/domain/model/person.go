package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Person struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name,omitempty" json:"name"`
}
