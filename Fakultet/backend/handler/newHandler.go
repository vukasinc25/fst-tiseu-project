package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"mime"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
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

func (nh *newHandler) GetAllRegistrationsToCompetition(w http.ResponseWriter, req *http.Request) {
	log.Println("Usli u GetAllRegistrationsToCompetition")

	vars := mux.Vars(req)
	competitionId := vars["id"]
	competitionId = strings.Trim(competitionId, "\"")
	log.Println("CompetitionId: ", competitionId)

	results, err := nh.repo.GetAllRegistrationsToCompetition(competitionId)
	if err != nil {
		log.Println("Error in GetAllRegistrationsToCompetition method: ", err)
		sendErrorWithMessage(w, "Fetch error", http.StatusInternalServerError)
		return
	}

	encodeToJson(w, results)
}

func (nh *newHandler) DiplomaRequest(w http.ResponseWriter, req *http.Request) { //ovde
	log.Println("Usli u DiplomaRequest")

	// authPayload, ok := req.Context().Value("authorization_payload").(*token.Payload)
	// if !ok {
	// 	// Handle case where authorization_payload is not found in context
	// 	http.Error(w, "authorization_payload not found in context", http.StatusInternalServerError)
	// 	return
	// }
	// log.Println("Payload: ", authPayload)

	// userId := authPayload.ID.Hex()
	// userId = strings.Trim(userId, "\"")
	// log.Println("User Id: ", userId)

	// diplomaRequest := model.DiplomaRequest{
	// 	ID:         primitive.NewObjectID(),
	// 	UserId:     userId,
	// 	IssueDate:  time.Now(),
	// 	InPending:  true,
	// 	IsApproved: false,
	// }
	vars := mux.Vars(req)

	id := vars["id"]
	name := vars["name"]

	id = strings.Trim(id, "\"")
	name = strings.Trim(name, "\"")

	log.Println("User Id: ", id)
	log.Println("User Name: ", name)

	diplomaRequest := model.DiplomaRequest{
		ID:         primitive.NewObjectID(),
		UserId:     id,
		UserName:   name,
		IssueDate:  time.Now(),
		InPending:  true,
		IsApproved: false,
	}

	err := nh.repo.CreateDiplomaRequest(&diplomaRequest)
	if err != nil {
		log.Println("Error: ", err)
		sendErrorWithMessage(w, "Cant send request", http.StatusInternalServerError)
		return
	}

	sendErrorWithMessage(w, "Sent", http.StatusOK)
}

func (nh *newHandler) GetDiplomaRequestInPendingState(w http.ResponseWriter, req *http.Request) {
	diplomas, err := nh.repo.GetAllDiplomaRequestsInPendingState()
	if err != nil {
		log.Println("Error: ", err)
		sendErrorWithMessage(w, "Cant get diploma requests", http.StatusInternalServerError)
		return
	}

	encodeToJson(w, diplomas)
}

func (nh *newHandler) GetDiplomaRequestsForUserId(w http.ResponseWriter, req *http.Request) { //ovde
	log.Println("Usli u GetDiplomaRequestsForUserId")
	// authPayload, ok := req.Context().Value("authorization_payload").(*token.Payload)
	// if !ok {
	// 	// Handle case where authorization_payload is not found in context
	// 	http.Error(w, "authorization_payload not found in context", http.StatusInternalServerError)
	// 	return
	// }
	// log.Println("Payload: ", authPayload)

	// userId := authPayload.ID.Hex()
	// userId = strings.Trim(userId, "\"")
	// log.Println("User Id: ", userId)

	vars := mux.Vars(req)

	userId := vars["id"]

	userId = strings.Trim(userId, "\"")

	requests, err := nh.repo.GetDiplomaRequestsForUserId(userId)
	if err != nil {
		log.Println("Cant get requests: ", err)
		sendErrorWithMessage(w, "Cant get requests", http.StatusInternalServerError)
		return
	}

	encodeToJson(w, requests)
}

func (nh *newHandler) DecideDiplomaRequest(w http.ResponseWriter, req *http.Request) {
	log.Println("Usli u DecideDiplomaRequest")
	vars := mux.Vars(req)

	id := vars["id"]

	id = strings.Trim(id, "\"")
	log.Println("Id: ", id)
	// var result = true

	isApproved, err := decodeIsApproved(req.Body)
	if err != nil {
		log.Println("Cant decode body: ", err)
		sendErrorWithMessage(w, "Cant decode body", http.StatusBadRequest)
		return
	}

	log.Println(isApproved.IsApproved)

	err = nh.repo.UpdateDiplomaRequest(id, isApproved.IsApproved)
	if err != nil {
		log.Println("Cant update diploma request: ", err)
		sendErrorWithMessage(w, "Cant update diploma request", http.StatusInternalServerError)
		return
	}

	diplomaRequest, err := nh.repo.GetDiplomaRequestById(id)
	if err != nil {
		log.Print("Cant find diploma request by id: ", err)
		sendErrorWithMessage(w, "Cant find diploma request by id", http.StatusInternalServerError)
		return
	}

	// create diplome
	if isApproved.IsApproved {
		err = nh.CreateDiploma(diplomaRequest.UserId, diplomaRequest.UserName)
		if err != nil {
			log.Print("Cant create diploma: ", err)
			sendErrorWithMessage(w, "Cant create diploma", http.StatusInternalServerError)
			return
		}
	}

	sendErrorWithMessage(w, "Ok", http.StatusOK)
}

func (nh *newHandler) CreateStudyProgram(w http.ResponseWriter, req *http.Request) {
	log.Println("Usli u CreateStudyProgram")

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
	rt, err := decodeStudyProgramBody(req.Body)
	if err != nil {
		log.Println("Decode: ", err)
		sendErrorWithMessage(w, "Error when decoding data", http.StatusBadRequest)
		return
	}

	rt.ID = primitive.NewObjectID()

	err = nh.repo.InsertStudyProgram(rt)
	if err != nil {
		log.Println(err)
		sendErrorWithMessage(w, "Cant create study program", http.StatusInternalServerError)
		return
	}

	sendErrorWithMessage(w, "Study progrma successfuly created", http.StatusCreated)
}

func (nh *newHandler) GetAlltudyPrograms(w http.ResponseWriter, req *http.Request) {
	log.Println("Usli u GetAllStudyPrograms")

	results, err := nh.repo.GetAllStudyPrograms()
	if err != nil {
		log.Println(err)
		sendErrorWithMessage(w, "Cant return study programs", http.StatusInternalServerError)
		return
	}

	encodeToJson(w, results)
}

func (nh *newHandler) GetAllDepartments(w http.ResponseWriter, req *http.Request) {
	log.Println("Usli u GetAllDepartments")

	departments, err := nh.repo.GetAllDepartments()
	if err != nil {
		log.Println("Cant get departments: ", err)
		sendErrorWithMessage(w, "Cant return departments", http.StatusInternalServerError)
		return
	}

	users, err := nh.repo.GetAllUsers()
	if err != nil {
		log.Println("Cant get users: ", err)
		sendErrorWithMessage(w, "Cant return users", http.StatusInternalServerError)
		return
	}

	var departments1 []model.Department
	for _, deptDB := range *departments {
		dept := model.Department{
			ID:    deptDB.ID,
			Name:  deptDB.Name,
			Staff: *users, // Add all users to the Staff field
		}
		departments1 = append(departments1, dept)
	}

	encodeToJson(w, departments1)

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

	log.Println("Before decodeBody")
	bodyBytes, _ := ioutil.ReadAll(req.Body)
	log.Println("Request Body: ", string(bodyBytes))
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	log.Println("Pre decodeBody")
	rt, err := decodeBody(req.Body)
	if err != nil {
		log.Println("Decode: ", err)
		sendErrorWithMessage(w, "Error when decoding data", http.StatusBadRequest)
		return
	}
	rt.ID = primitive.NewObjectID()

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

func (nh *newHandler) CreateDiploma(userId string, userName string) error {
	log.Println("Usli u CreateDiploma")

	// generate random number between 6 and 10
	randomNumber := rand.Intn(5) + 6
	log.Println("Rand number: ", randomNumber)

	diploma := model.Diploma{
		ID:           primitive.NewObjectID(),
		UserId:       userId,
		UserName:     userName,
		IssueDate:    time.Now(),
		AverageGrade: strconv.Itoa(randomNumber),
	}

	err := nh.repo.InsertDiploma(&diploma)
	if err != nil {
		log.Println("Cant create diploma in mongo: ", err)
		return err
	}

	return nil
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

	err = nh.repo.InsertCompetition(rt)
	if err != nil {
		log.Println(err)
		sendErrorWithMessage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendErrorWithMessage(w, "Competition Created", http.StatusCreated)
}

func (nh *newHandler) GetAllCompetitions(w http.ResponseWriter, req *http.Request) {
	log.Println("Usli u GetAllCompetitions")

	results, err := nh.repo.GetAllCompetitions()
	if err != nil {
		log.Println(err)
		sendErrorWithMessage(w, "Cant return competitions", http.StatusInternalServerError)
		return
	}

	encodeToJson(w, results)
}

func (nh *newHandler) GetCompetitionById(w http.ResponseWriter, req *http.Request) {
	log.Println("Usli u GetCompetitionById")

	vars := mux.Vars(req)

	id := vars["id"]

	id = strings.Trim(id, "\"")
	log.Println("Id: ", id)

	results, err := nh.repo.GetCompetitionById(id)
	if err != nil {
		log.Println(err)
		sendErrorWithMessage(w, "Cant return competition", http.StatusInternalServerError)
		return
	}

	encodeToJson(w, results)
}

// Registrating logged user to the Fakulty Competition
func (nh *newHandler) CreateRegistrationUserToCompetition(w http.ResponseWriter, req *http.Request) {
	log.Println("Usli u CreateRegistrationUserToCompetition")

	//TREBA TI ZA POVEZIVANJE SA SREDNOJM SKOLOM I DOBAVLJANJE DIPLOME OD NJE!!!!!
	// ctx := req.Context()

	// token, ok := ctx.Value("accessToken").(string)
	// if !ok || token == "" {
	// 	sendErrorWithMessage(w, "Authorization token not found", http.StatusInternalServerError)
	// 	return
	// }

	// log.Println("Token: ", token)

	// req.Header.Set("Content-Type", "application/json")

	// // Set the Authorization header with the Bearer token
	// req.Header.Set("Authorization", "Bearer "+token)

	// // send request to the high school service(high school service not created still)

	vars := mux.Vars(req)

	competitionId := vars["id"]
	userId := vars["userId"]
	userName := vars["userName"]

	competitionId = strings.Trim(competitionId, "\"")
	userId = strings.Trim(userId, "\"")
	userName = strings.Trim(userName, "\"")
	log.Println("CompetitionId: ", competitionId)

	url := "http://skola:8002/skola/diplomas"

	requestBody := map[string]string{
		"userId": userId, // Replace with the actual userId
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return
	}

	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	if resp == nil {
		sendErrorWithMessage(w, "User cant registerd to the competition", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	log.Println("Response:", string(body))
	log.Println("UserName:", userName)

	if body == nil {
		sendErrorWithMessage(w, "User cant registerd to the competition", http.StatusBadRequest)
		return
	}

	// log.Println("Body: ", string(body))

	// authPayload, ok := req.Context().Value("authorization_payload").(*token.Payload)
	// if !ok {
	// 	// Handle case where authorization_payload is not found in context
	// 	http.Error(w, "authorization_payload not found in context", http.StatusInternalServerError)
	// 	return
	// }
	// log.Println("Payload: ", authPayload)

	// userId := authPayload.ID.Hex()
	// userId = strings.Trim(userId, "\"")

	/////// OVDE
	userRegistration := model.RegisteredStudentToCommpetition{
		CompetitionID: competitionId,
		UserID:        userId,
		UserName:      userName,
	}

	userRegistration.ID = primitive.NewObjectID()

	err = nh.repo.CreateRegisteredStudentToTheCommpetition(&userRegistration)
	if err != nil {
		log.Println(err)
		sendErrorWithMessage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendErrorWithMessage(w, "User successfuly registerd to the competition", http.StatusCreated)
}

func (nh *newHandler) CreateUserExamResult(w http.ResponseWriter, req *http.Request) {
	log.Println("Usli u CreateUserExamResult")

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
	rt, err := decodeExamResultBody(req.Body)
	if err != nil {
		log.Println("Decode: ", err)
		sendErrorWithMessage(w, "Error when decoding data", http.StatusBadRequest)
		return
	}

	rt.ID = primitive.NewObjectID()
	rt.ScoreEntryDate = time.Now()

	err = nh.repo.InsertUserExamResult(rt)
	if err != nil {
		log.Println(err)
		sendErrorWithMessage(w, "Cant insert user exam result", http.StatusInternalServerError)
		return
	}

	sendErrorWithMessage(w, "User exam result inserted", http.StatusCreated)
}

func (nh *newHandler) GetDiplomaByUserId(w http.ResponseWriter, req *http.Request) { //ovde od 508-518
	log.Println("Usli u GetDiplomaByUserId")
	//treba ti payload zato sto nemas id od logovanog korisnika tako da kada posalje zahtev ne moras da prosledjujes id

	// authPayload, ok := req.Context().Value("authorization_payload").(*token.Payload)
	// if !ok {
	// 	// Handle case where authorization_payload is not found in context
	// 	http.Error(w, "authorization_payload not found in context", http.StatusInternalServerError)
	// 	return
	// }
	// log.Println("Payload: ", authPayload)

	// id := authPayload.ID.Hex()
	// id = strings.Trim(id, "\"")
	// log.Println("Id: ", id)

	vars := mux.Vars(req)

	id := vars["id"]

	id = strings.Trim(id, "\"")

	// vars := mux.Vars(req)

	// id := vars["id"]

	// id = strings.Trim(id, "\"")
	// log.Println("Id: ", id)

	diploma, err := nh.repo.GetDiplomaByUserId(id)
	if err != nil {
		log.Println(err)
		sendErrorWithMessage(w, "User with that id has no diploma", http.StatusInternalServerError)
		return
	}

	encodeToJson(w, diploma)
}

func (nh *newHandler) GetStudyProgramById(w http.ResponseWriter, req *http.Request) {
	log.Println("Usli u GetStudyProgramById")

	vars := mux.Vars(req)

	id := vars["id"]

	id = strings.Trim(id, "\"")
	log.Println("Id: ", id)

	studyProgram, err := nh.repo.GetStudyProgramId(id)
	if err != nil {
		log.Println(err)
		sendErrorWithMessage(w, "Study program with that id dont exist", http.StatusInternalServerError)
		return
	}

	encodeToJson(w, studyProgram)
}

func (nh *newHandler) GetAllExamResultsByCompetitionId(w http.ResponseWriter, req *http.Request) {
	log.Println("Usli u GetAllExamResultsByCompetitionId")
	vars := mux.Vars(req)

	id := vars["id"]

	id = strings.Trim(id, "\"")
	log.Println("Id: ", id)

	results, err := nh.repo.GetAllExamResultsByCompetitionId(id)
	if err != nil {
		log.Println(err)
		sendErrorWithMessage(w, "User with that id has no diploma", http.StatusInternalServerError)
		return
	}

	encodeToJson(w, results)
}

func (nh *newHandler) CreateDepartment(w http.ResponseWriter, req *http.Request) {
	log.Println("Usli u CreateDepartment")

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
	rt, err := decodeDepartmentBody(req.Body)
	if err != nil {
		log.Println("Decode: ", err)
		sendErrorWithMessage(w, "Error when decoding data", http.StatusBadRequest)
		return
	}

	rt.ID = primitive.NewObjectID()

	err = nh.repo.InsertDepartment(rt)
	if err != nil {
		log.Println(err)
		sendErrorWithMessage(w, "Cant insert department", http.StatusInternalServerError)
		return
	}

	sendErrorWithMessage(w, "Department inserted", http.StatusCreated)
}

func decodeIsApproved(r io.Reader) (*model.IsApproved, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt model.IsApproved
	if err := dec.Decode(&rt); err != nil {
		log.Println("Decode cant be done")
		return nil, err
	}

	return &rt, nil
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

func decodeExamResultBody(r io.Reader) (*model.ExamResult, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt model.ExamResult
	if err := dec.Decode(&rt); err != nil {
		log.Println("Decode cant be done")
		return nil, err
	}

	return &rt, nil
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

func decodeStudyProgramBody(r io.Reader) (*model.StudyProgram, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt model.StudyProgram
	if err := dec.Decode(&rt); err != nil {
		log.Println("Decode cant be done")
		return nil, err
	}

	return &rt, nil
}

func decodeDepartmentBody(r io.Reader) (*model.DepartmentDB, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt model.DepartmentDB
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

func sendErrorWithMessage(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorResponse := map[string]string{"message": message}
	json.NewEncoder(w).Encode(errorResponse)
}

func encodeToJson(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Error encoding departments to JSON:", err)
		sendErrorWithMessage(w, "Error encoding response", http.StatusInternalServerError)
	}
}
