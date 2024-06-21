package repository

import (
	"context"
	"time"

	"github.com/MariliaNeves/api-genealogy/internal/domain/model"
	"github.com/MariliaNeves/api-genealogy/internal/domain/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_PERSON_COLLECTION = "MONGODB_PERSON_COLLECTION"
)

type personRepository struct {
	collection *mongo.Collection
}

func NewPersonRepository(db *mongo.Database) repository.PersonRepository {
	return &personRepository{collection: db.Collection(MONGODB_PERSON_COLLECTION)}
}

func (r *personRepository) Create(person model.Person) (model.Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := r.collection.InsertOne(ctx, person)
	if err != nil {
		return person, err
	}
	person.ID = result.InsertedID.(primitive.ObjectID)
	return person, nil
}

func (r *personRepository) GetAll() ([]model.Person, error) {
	var people []model.Person
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return people, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person model.Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	return people, nil
}

func (r *personRepository) GetByID(id string) (model.Person, error) {
	var person model.Person
	objID, _ := primitive.ObjectIDFromHex(id)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&person)
	return person, err
}

func (r *personRepository) Update(id string, person model.Person) (model.Person, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.D{
		{"$set", person},
	})
	return person, err
}

func (r *personRepository) Delete(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
