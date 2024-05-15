package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime"
	"net/http"

	"github.com/vukasinc25/fst-tiseu-project/model"
	"github.com/vukasinc25/fst-tiseu-project/repository"
)

type newHandler struct {
	repo *repository.NewRepository
}

func NewHandler(r *repository.NewRepository) (*newHandler, error) {
	return &newHandler{r}, nil
}

func (nh *newHandler) CreateCompetition(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		log.Println("Error cant mimi.ParseMediaType")
		sendErrorWithMessage(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("expect application/json Content-Type")
		sendErrorWithMessage(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeCompetitionBody(req.Body)
	if err != nil {
		log.Println("Decode: ", err)
		sendErrorWithMessage(w, "Error when decoding data", http.StatusBadRequest)
		return
	}

	log.Println("User: ", rt)

	err = nh.repo.CreateCompetition(rt)
	if err != nil {
		log.Println(err)
		sendErrorWithMessage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendErrorWithMessage(w, "User Created", http.StatusCreated)
}

func decodeCompetitionBody(r io.Reader) (*model.Competition, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt model.Competition
	if err := dec.Decode(&rt); err != nil {
		log.Println("Decode cant be done")
		return nil, err
	}

	return &rt, nil
}

func sendErrorWithMessage(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorResponse := map[string]string{"message": message}
	json.NewEncoder(w).Encode(errorResponse)
}
