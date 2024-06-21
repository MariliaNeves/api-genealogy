package repository

import "github.com/MariliaNeves/api-genealogy/internal/domain/model"

type RelationshipRepository interface {
	Create(relationship model.Relationship) (model.Relationship, error)
	GetAll() ([]model.Relationship, error)
	GetByID(id string) (model.Relationship, error)
	Update(id string, relationship model.Relationship) (model.Relationship, error)
	Delete(id string) error
}
