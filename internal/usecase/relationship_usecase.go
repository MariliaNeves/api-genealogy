package usecase

import (
	"github.com/MariliaNeves/api-genealogy/internal/domain/model"
	"github.com/MariliaNeves/api-genealogy/internal/domain/repository"
)

type RelationshipUsecase interface {
	CreateRelationship(relationship model.Relationship) (model.Relationship, error)
	GetPeople() ([]model.Relationship, error)
	GetRelationship(id string) (model.Relationship, error)
	UpdateRelationship(id string, relationship model.Relationship) (model.Relationship, error)
	DeleteRelationship(id string) error
}

type relationshipUsecase struct {
	relationshipRepo repository.RelationshipRepository
}

func NewRelationshipUsecase(repo repository.RelationshipRepository) RelationshipUsecase {
	return &relationshipUsecase{relationshipRepo: repo}
}

func (u *relationshipUsecase) CreateRelationship(relationship model.Relationship) (model.Relationship, error) {
	return u.relationshipRepo.Create(relationship)
}

func (u *relationshipUsecase) GetPeople() ([]model.Relationship, error) {
	return u.relationshipRepo.GetAll()
}

func (u *relationshipUsecase) GetRelationship(id string) (model.Relationship, error) {
	return u.relationshipRepo.GetByID(id)
}

func (u *relationshipUsecase) UpdateRelationship(id string, relationship model.Relationship) (model.Relationship, error) {
	return u.relationshipRepo.Update(id, relationship)
}

func (u *relationshipUsecase) DeleteRelationship(id string) error {
	return u.relationshipRepo.Delete(id)
}
