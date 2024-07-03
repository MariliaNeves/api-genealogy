package main

import (
	"context"
	"log"
	"net/http"

	"github.com/MariliaNeves/api-genealogy/internal/adapters/api"
	"github.com/MariliaNeves/api-genealogy/internal/adapters/repository/mongodb"
	"github.com/MariliaNeves/api-genealogy/internal/usecase"
	"github.com/MariliaNeves/api-genealogy/server/config"
	"github.com/MariliaNeves/api-genealogy/server/config/logger"
)

func main() {
	logger.Info("Start App")

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
