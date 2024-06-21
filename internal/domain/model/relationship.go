package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Relationship struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Parent   primitive.ObjectID `bson:"parent,omitempty"`
	Children primitive.ObjectID `bson:"children,omitempty"`
}
