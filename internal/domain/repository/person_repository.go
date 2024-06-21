package repository

import "github.com/MariliaNeves/api-genealogy/internal/domain/model"

type PersonRepository interface {
	Create(person model.Person) (model.Person, error)
	GetAll() ([]model.Person, error)
	GetByID(id string) (model.Person, error)
	Update(id string, person model.Person) (model.Person, error)
	Delete(id string) error
}
