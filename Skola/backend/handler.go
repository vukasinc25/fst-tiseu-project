package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	db *Repo
}

func NewHandler(repo *Repo) (*Handler, error) {
	return &Handler{db: repo}, nil
}

func (hd *Handler) GetUserDiplomas(writer http.ResponseWriter, request *http.Request) {

	id := &RequestBody{}
	d := json.NewDecoder(request.Body)
	err := d.Decode(&id)
	if err != nil {
		log.Println("Error decoding request body")
	}

	diplomas, err := hd.db.GetDiplomaByStudent(id.UserId)
	if err != nil {
		log.Println(err)
		http.Error(writer, "Error in db", http.StatusNotFound)
	}

	e := json.NewEncoder(writer)
	err = e.Encode(diplomas)
	if err != nil {
		http.Error(writer, "Unable to convert to json", http.StatusInternalServerError)
	}
}
