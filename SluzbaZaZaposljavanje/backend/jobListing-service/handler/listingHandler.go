package handler

import (
	"encoding/json"
	"errors"
	"github.com/vukasinc25/fst-tiseu-project/repository"
	"log"
	"mime"
	"net/http"
)

type ListingHandler struct {
	logger *log.Logger
	repo   *repository.ListingRepository
}

func NewHandler(l *log.Logger, r *repository.ListingRepository) (*ListingHandler, error) {
	return &ListingHandler{l, r}, nil
}

func (*ListingHandler) CreateJobListing(w http.ResponseWriter, req *http.Request) {
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

	sendErrorWithMessage(w, "User Created", http.StatusCreated)
}

func sendErrorWithMessage(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorResponse := map[string]string{"message": message}
	json.NewEncoder(w).Encode(errorResponse)
}
