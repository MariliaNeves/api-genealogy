package main

import (
	"context"
	"log"
	"myapp/internal/adapters/api"
	"myapp/internal/adapters/repository/mongo"
	"myapp/internal/usecase"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	db := client.Database("myappdb")

	personRepo := mongo.NewPersonRepository(db)
	relationshipRepo := mongo.NewRelationshipRepository(db)

	personUsecase := usecase.NewPersonUsecase(personRepo)
	relationshipUsecase := usecase.NewRelationshipUsecase(relationshipRepo)

	router := api.NewRouter(personUsecase, relationshipUsecase)

	log.Fatal(http.ListenAndServe(":8000", router))
}
