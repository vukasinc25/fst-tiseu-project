package handler

import (
	"encoding/json"
	"errors"
	"github.com/vukasinc25/fst-tiseu-project/model"
	"github.com/vukasinc25/fst-tiseu-project/repository"
	"io"
	"log"
	"mime"
	"net/http"
	"strings"
)

type ListingHandler struct {
	logger *log.Logger
	repo   *repository.ListingRepository
}

func NewHandler(l *log.Logger, r *repository.ListingRepository) (*ListingHandler, error) {
	return &ListingHandler{l, r}, nil
}

type KeyProduct struct{}

func (lh *ListingHandler) CreateJobListing(w http.ResponseWriter, req *http.Request) {
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

	jobListing, err := decodeBody(req.Body)

	err = lh.repo.Insert(jobListing)
	if err != nil {
		lh.logger.Println("error:1", err.Error())
		if strings.Contains(err.Error(), "duplicate key") {
			sendErrorWithMessage(w, "accommodation with that name already exists", http.StatusBadRequest)
			return
		}
		sendErrorWithMessage(w, "NE VALJA", http.StatusBadRequest)
		return
	}

	sendErrorWithMessage(w, "BAS NE VALJA", http.StatusCreated)
}

func decodeBody(r io.Reader) (*model.JobListing, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt model.JobListing
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
