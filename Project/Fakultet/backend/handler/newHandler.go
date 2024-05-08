package handler

import (
	"encoding/json"
	"errors"
	"log"
	"mime"
	"net/http"

	"github.com/vukasinc25/fst-tiseu-projet/repository"
)

type newHandler struct {
	repo *repository.NewRepository
}

func NewHandler(r *repository.NewRepository) (*newHandler, error) {
	return &newHandler{r}, nil
}

func (*newHandler) CreateUser(w http.ResponseWriter, req *http.Request) {
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

// func decodeBody(r io.Reader) (*model.User, error) {
// 	dec := json.NewDecoder(r)
// 	dec.DisallowUnknownFields()

// 	var rt model.User
// 	if err := dec.Decode(&rt); err != nil {
// 		log.Println("Decode cant be done")
// 		return nil, err
// 	}

// 	return &rt, nil
// }

func sendErrorWithMessage(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorResponse := map[string]string{"message": message}
	json.NewEncoder(w).Encode(errorResponse)
}
