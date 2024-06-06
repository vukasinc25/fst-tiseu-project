package handler

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/vukasinc25/fst-tiseu-project/model"
	"github.com/vukasinc25/fst-tiseu-project/repository"
	"io"
	"log"
	"mime"
	"net/http"
	"strings"
)

type JobHandler struct {
	logger *log.Logger
	repo   *repository.JobRepository
}

func NewHandler(l *log.Logger, r *repository.JobRepository) (*JobHandler, error) {
	return &JobHandler{l, r}, nil
}

func (jh *JobHandler) GetAllJobListings(w http.ResponseWriter, req *http.Request) {
	jobListings, err := jh.repo.GetAllJobListings()
	if err != nil {
		sendErrorWithMessage(w, err.Error(), http.StatusUnsupportedMediaType)
	}

	err = jobListings.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to convert to json", http.StatusInternalServerError)
		return
	}
}

func (jh *JobHandler) GetJobListing(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	jobListing, err := jh.repo.GetJobListing(id)
	if err != nil {
		sendErrorWithMessage(w, err.Error(), http.StatusUnsupportedMediaType)
	}

	if &jobListing == nil {
		sendErrorWithMessage(w, err.Error(), http.StatusBadRequest)
	}

	err = jobListing.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to convert to json", http.StatusInternalServerError)
		return
	}
}

func (jh *JobHandler) CreateJobListing(w http.ResponseWriter, req *http.Request) {
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

	jobListing, err := decodeListingBody(req.Body)

	err = jh.repo.InsertJobListing(jobListing)
	if err != nil {
		jh.logger.Println("error:1", err.Error())
		if strings.Contains(err.Error(), "duplicate key") {
			sendErrorWithMessage(w, "accommodation with that name already exists", http.StatusBadRequest)
			return
		}
		sendErrorWithMessage(w, "Error while inserting job listing", http.StatusBadRequest)
		return
	}
}

func (jh *JobHandler) GetAllJobApplicationsByEmployerId(w http.ResponseWriter, req *http.Request) {
	//ctx := req.Context()
	//
	//authPayload, ok := ctx.Value("authorization_payload").(*token.Payload)
	//if !ok || authPayload == nil {
	//	sendErrorWithMessage(w, "Authorization payload not found", http.StatusInternalServerError)
	//	return
	//}
	//log.Println("Payload: ", authPayload)
	//
	//id := authPayload.ID.Hex()
	//id = strings.Trim(id, "\"")
	//log.Println("Id: ", id)
	vars := mux.Vars(req)
	id := vars["id"]

	jobListing, err := jh.repo.GetAllJobApplicationsByEmployerId(id)
	if err != nil {
		sendErrorWithMessage(w, err.Error(), http.StatusUnsupportedMediaType)
	}

	if &jobListing == nil {
		sendErrorWithMessage(w, err.Error(), http.StatusBadRequest)
	}

	err = jobListing.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to convert to json", http.StatusInternalServerError)
		return
	}
}

func (jh *JobHandler) CreateJobApplication(w http.ResponseWriter, req *http.Request) {
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

	jobApplication, err := decodeApplicationBody(req.Body)

	//diploma, err := getDiplomaFromFakultetService()
	//log.Println(diploma)

	err = jh.repo.InsertJobApplication(jobApplication)
	if err != nil {
		jh.logger.Println("error:1", err.Error())
		if strings.Contains(err.Error(), "duplicate key") {
			sendErrorWithMessage(w, "accommodation with that name already exists", http.StatusBadRequest)
			return
		}
		sendErrorWithMessage(w, "Error while inserting job application", http.StatusBadRequest)
		return
	}
}

// COMMUNICATION WITH FAKULTET SERVICE
func getDiplomaFromFakultetService(token string) (*http.Response, error) {
	url := "http://auth-service:8001/fakultet/user/diplomaByUserId"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set Authorization header with bearer token
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	httpResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return httpResp, nil
}

func decodeListingBody(r io.Reader) (*model.JobListing, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt model.JobListing
	if err := dec.Decode(&rt); err != nil {
		log.Println("Decode cant be done")
		return nil, err
	}

	return &rt, nil
}

func decodeApplicationBody(r io.Reader) (*model.JobApplication, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt model.JobApplication
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
