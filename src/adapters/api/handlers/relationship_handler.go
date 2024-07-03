package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MariliaNeves/api-genealogy/src/domain/model"
	"github.com/MariliaNeves/api-genealogy/src/usecase"

	"github.com/gorilla/mux"
)

type RelationshipHandler struct {
	usecase usecase.RelationshipUsecase
}

func NewRelationshipHandler(router *mux.Router, usecase usecase.RelationshipUsecase) {
	handler := &RelationshipHandler{usecase: usecase}
	router.HandleFunc("/relationship", handler.CreateRelationship).Methods("POST")
	router.HandleFunc("/relationships", handler.GetRelationships).Methods("GET")
	router.HandleFunc("/relationship/{id}", handler.GetRelationship).Methods("GET")
	router.HandleFunc("/relationship/{id}", handler.UpdateRelationship).Methods("PUT")
	router.HandleFunc("/relationship/{id}", handler.DeleteRelationship).Methods("DELETE")
}

func (h *RelationshipHandler) CreateRelationship(w http.ResponseWriter, r *http.Request) {
	var person model.Relationship
	_ = json.NewDecoder(r.Body).Decode(&person)
	createdRelationship, err := h.usecase.CreateRelationship(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdRelationship)
}

func (h *RelationshipHandler) GetRelationships(w http.ResponseWriter, r *http.Request) {
	relationship, err := h.usecase.GetRelationships()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(relationship)
}

func (h *RelationshipHandler) GetRelationship(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	person, err := h.usecase.GetRelationship(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(person)
}

func (h *RelationshipHandler) UpdateRelationship(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person model.Relationship
	_ = json.NewDecoder(r.Body).Decode(&person)
	updatedRelationship, err := h.usecase.UpdateRelationship(params["id"], person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedRelationship)
}

func (h *RelationshipHandler) DeleteRelationship(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := h.usecase.DeleteRelationship(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Relationship deleted")
}
