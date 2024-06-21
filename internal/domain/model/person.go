package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Person struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name,omitempty"`
	RelationshipID primitive.ObjectID `bson:"relationship,omitempty"`
}
