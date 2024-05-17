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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type newHandler struct {
	repo *repository.NewRepository
}

func NewHandler(r *repository.NewRepository) (*newHandler, error) {
	return &newHandler{r}, nil
}

func (nh *newHandler) CreateFirm(w http.ResponseWriter, req *http.Request) {
	log.Println("Usli u Create")
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

	rt, err := decodeFirmBody(req.Body)
	if err != nil {
		log.Println("Decode: ", err)
		sendErrorWithMessage(w, "Error when decoding data", http.StatusBadRequest)
		return
	}

	rt.ID = primitive.NewObjectID()

	log.Println("Firm: ", rt)

	err = nh.repo.CreateFirm(rt)
	if err != nil {
		log.Println(err)
		sendErrorWithMessage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendErrorWithMessage(w, "Firm Created", http.StatusCreated)
}

func decodeFirmBody(r io.Reader) (*model.Firm, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt model.Firm
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
