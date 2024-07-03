package main

import (
	"context"
	"log"
	"net/http"

	"github.com/MariliaNeves/api-genealogy/config"
	"github.com/MariliaNeves/api-genealogy/config/logger"
	"github.com/MariliaNeves/api-genealogy/src/adapters/api"
	"github.com/MariliaNeves/api-genealogy/src/adapters/repository/mongodb"
	"github.com/MariliaNeves/api-genealogy/src/usecase"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Start App")
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	db, err := config.NewMongoDBConnection(context.Background())
	if err != nil {
		logger.Error("Error trying to connect to database", err)
	}

	personRepo := mongodb.NewPersonRepository(db)
	relationshipRepo := mongodb.NewRelationshipRepository(db)

	personUsecase := usecase.NewPersonUsecase(personRepo)
	relationshipUsecase := usecase.NewRelationshipUsecase(relationshipRepo)

	router := api.NewRouter(personUsecase, relationshipUsecase)

	log.Fatal(http.ListenAndServe(":8000", router))
}
