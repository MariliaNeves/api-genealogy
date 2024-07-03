package mongodb

import (
	"context"
	"os"
	"time"

	"github.com/MariliaNeves/api-genealogy/internal/domain/model"
	"github.com/MariliaNeves/api-genealogy/internal/domain/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_RELATIONSHIP_COLLECTION = "MONGODB_RELATIONSHIP_COLLECTION"
)

type relationshipRepository struct {
	collection *mongo.Collection
}

func NewRelationshipRepository(db *mongo.Database) repository.RelationshipRepository {
	collection_name := os.Getenv(MONGODB_RELATIONSHIP_COLLECTION)
	return &relationshipRepository{collection: db.Collection(collection_name)}
}

func (r *relationshipRepository) Create(relationship model.Relationship) (model.Relationship, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := r.collection.InsertOne(ctx, relationship)
	if err != nil {
		return relationship, err
	}
	relationship.ID = result.InsertedID.(primitive.ObjectID)
	return relationship, nil
}

func (r *relationshipRepository) GetAll() ([]model.Relationship, error) {
	var people []model.Relationship
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return people, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var relationship model.Relationship
		cursor.Decode(&relationship)
		people = append(people, relationship)
	}
	return people, nil
}

func (r *relationshipRepository) GetByID(id string) (model.Relationship, error) {
	var relationship model.Relationship
	objID, _ := primitive.ObjectIDFromHex(id)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&relationship)
	return relationship, err
}

func (r *relationshipRepository) Update(id string, relationship model.Relationship) (model.Relationship, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.D{
		{"$set", relationship},
	})
	return relationship, err
}

func (r *relationshipRepository) Delete(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
