package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MariliaNeves/api-genealogy/internal/domain/model"
	"github.com/MariliaNeves/api-genealogy/internal/usecase"

	"github.com/gorilla/mux"
)

type PersonHandler struct {
	usecase usecase.PersonUsecase
}

func NewPersonHandler(router *mux.Router, usecase usecase.PersonUsecase) {
	handler := &PersonHandler{usecase: usecase}
	router.HandleFunc("/people", handler.CreatePerson).Methods("POST")
	router.HandleFunc("/people", handler.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", handler.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", handler.UpdatePerson).Methods("PUT")
	router.HandleFunc("/people/{id}", handler.DeletePerson).Methods("DELETE")
}

func (h *PersonHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person model.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	createdPerson, err := h.usecase.CreatePerson(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdPerson)
}

func (h *PersonHandler) GetPeople(w http.ResponseWriter, r *http.Request) {
	people, err := h.usecase.GetPeople()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(people)
}

func (h *PersonHandler) GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	person, err := h.usecase.GetPerson(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(person)
}

func (h *PersonHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person model.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	updatedPerson, err := h.usecase.UpdatePerson(params["id"], person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedPerson)
}

func (h *PersonHandler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := h.usecase.DeletePerson(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Person deleted")
}
