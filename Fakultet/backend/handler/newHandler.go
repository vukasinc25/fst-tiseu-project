package handler

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime"
	"net/http"
	"strings"
	"time"

	"github.com/vukasinc25/fst-tiseu-project/model"
	"github.com/vukasinc25/fst-tiseu-project/repository"
	"github.com/vukasinc25/fst-tiseu-project/token"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type newHandler struct {
	repo *repository.NewRepository
}

func NewHandler(r *repository.NewRepository) (*newHandler, error) {
	return &newHandler{r}, nil
}

func (nh *newHandler) CreateUser(w http.ResponseWriter, req *http.Request) {
	log.Println("Usli u Create")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	log.Println("Pre decodeBody")
	rt, err := decodeBody(req.Body)
	if err != nil {
		log.Println("Decode: ", err)
		sendErrorWithMessage(w, "Error when decoding data", http.StatusBadRequest)
		return
	}

	err = nh.repo.Insert(rt, ctx)
	if err != nil {
		if strings.Contains(err.Error(), "username") {
			sendErrorWithMessage(w, "Provide different username", http.StatusConflict)
		} else if strings.Contains(err.Error(), "email") {
			sendErrorWithMessage(w, "Provide different email", http.StatusConflict)
		}
		return
	}

	sendErrorWithMessage(w, "User Created", http.StatusCreated)
}

func (nh *newHandler) CreateDiploma(w http.ResponseWriter, req *http.Request) {
	log.Println("Usli u CreateDiploma")

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

	log.Println("Pre decodeBody")
	rt, err := decodeDipomaBody(req.Body)
	if err != nil {
		log.Println("Decode: ", err)
		sendErrorWithMessage(w, "Error when decoding data", http.StatusBadRequest)
		return
	}

	rt.ID = primitive.NewObjectID()
	rt.IssueDate = time.Now()

	err = nh.repo.InsertDiploma(rt)
	if err != nil {
		sendErrorWithMessage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendErrorWithMessage(w, "User Diploma", http.StatusCreated)
}

func (nh *newHandler) CreateCompetition(w http.ResponseWriter, req *http.Request) {
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

	rt, err := decodeCompetitionBody(req.Body)
	if err != nil {
		log.Println("Decode: ", err)
		sendErrorWithMessage(w, "Error when decoding data", http.StatusBadRequest)
		return
	}

	rt.ID = primitive.NewObjectID()

	log.Println("Competition: ", rt)

	err = nh.repo.CreateCompetition(rt)
	if err != nil {
		log.Println(err)
		sendErrorWithMessage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendErrorWithMessage(w, "Competition Created", http.StatusCreated)
}

func (nh *newHandler) CreateRegistrationUserToCompetition(w http.ResponseWriter, req *http.Request) {
	log.Println("Usli u CreateRegistrationUserToCompetition")
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

	rt, err := decodeRegistrationOfUserToCompetitionBody(req.Body)
	if err != nil {
		log.Println("Decode: ", err)
		sendErrorWithMessage(w, "Error when decoding data", http.StatusBadRequest)
		return
	}

	rt.ID = primitive.NewObjectID()

	log.Println("Competition: ", rt)

	err = nh.repo.CreateRegisteredStudentToTheCommpetition(rt)
	if err != nil {
		log.Println(err)
		sendErrorWithMessage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendErrorWithMessage(w, "User successfuly registerd to the competition", http.StatusCreated)
}

func (nh *newHandler) GetDiplomaByUserId(w http.ResponseWriter, req *http.Request) {
	log.Println("Usli u GetDiplomaByUserId")

	ctx := req.Context()

	authPayload, ok := ctx.Value("authorization_payload").(*token.Payload)
	if !ok || authPayload == nil {
		sendErrorWithMessage(w, "Authorization payload not found", http.StatusInternalServerError)
		return
	}

	log.Println("Payload: ", authPayload)

	id := "6646761f9e10566e77913d79"

	diploma, err := nh.repo.GetDiplomaByUserId(id)
	if err != nil {
		log.Println(err)
		sendErrorWithMessage(w, "User with that id has no diploma", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(diploma); err != nil {
		log.Println("Error encoding diploma to JSON:", err)
		sendErrorWithMessage(w, "Error encoding response", http.StatusInternalServerError)
	}

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

func decodeRegistrationOfUserToCompetitionBody(r io.Reader) (*model.RegisteredStudentsToCommpetition, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt model.RegisteredStudentsToCommpetition
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

func decodeDipomaBody(r io.Reader) (*model.Diploma, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt model.Diploma
	if err := dec.Decode(&rt); err != nil {
		log.Println("Decode cant be done")
		return nil, err
	}

	return &rt, nil
}
func decodeBody(r io.Reader) (*model.User, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt model.User
	if err := dec.Decode(&rt); err != nil {
		log.Println("Decode cant be done")
		return nil, err
	}

	return &rt, nil
}
