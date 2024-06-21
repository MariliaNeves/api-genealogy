package usecase

import (
	"github.com/MariliaNeves/api-genealogy/internal/domain/model"
	"github.com/MariliaNeves/api-genealogy/internal/domain/repository"
)

type PersonUsecase interface {
	CreatePerson(person model.Person) (model.Person, error)
	GetPeople() ([]model.Person, error)
	GetPerson(id string) (model.Person, error)
	UpdatePerson(id string, person model.Person) (model.Person, error)
	DeletePerson(id string) error
}

type personUsecase struct {
	personRepo repository.PersonRepository
}

func NewPersonUsecase(repo repository.PersonRepository) PersonUsecase {
	return &personUsecase{personRepo: repo}
}

func (u *personUsecase) CreatePerson(person model.Person) (model.Person, error) {
	return u.personRepo.Create(person)
}

func (u *personUsecase) GetPeople() ([]model.Person, error) {
	return u.personRepo.GetAll()
}

func (u *personUsecase) GetPerson(id string) (model.Person, error) {
	return u.personRepo.GetByID(id)
}

func (u *personUsecase) UpdatePerson(id string, person model.Person) (model.Person, error) {
	return u.personRepo.Update(id, person)
}

func (u *personUsecase) DeletePerson(id string) error {
	return u.personRepo.Delete(id)
}
