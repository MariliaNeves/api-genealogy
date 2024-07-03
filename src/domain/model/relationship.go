package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Relationship struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Parent      primitive.ObjectID   `bson:"parentID,omitempty" json:"parentID"`
	ChildrenIDs []primitive.ObjectID `bson:"childrenIDs,omitempty" json:"childrenIDs"`
}
