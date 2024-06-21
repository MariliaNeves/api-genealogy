package api

import (
	"github.com/MariliaNeves/api-genealogy/internal/adapters/api/handlers"
	"github.com/MariliaNeves/api-genealogy/internal/usecase"

	"github.com/gorilla/mux"
)

func NewRouter(personUsecase usecase.PersonUsecase, relationshipUsecase usecase.RelationshipUsecase) *mux.Router {
	router := mux.NewRouter()

	handlers.NewPersonHandler(router, personUsecase)
	handlers.NewRelationshipHandler(router, relationshipUsecase)

	return router
}
