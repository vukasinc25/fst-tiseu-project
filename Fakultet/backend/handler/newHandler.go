package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
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

	sendErrorWithMessage(w, "User diploma created", http.StatusCreated)
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

	ctx := req.Context()

	token, ok := ctx.Value("accessToken").(string)
	if !ok || token == "" {
		sendErrorWithMessage(w, "Authorization token not found", http.StatusInternalServerError)
		return
	}

	log.Println("Token: ", token)

	req.Header.Set("Content-Type", "application/json")

	// Set the Authorization header with the Bearer token
	req.Header.Set("Authorization", "Bearer "+token)

	// send request to the high school service(high school service not created still)
	url := "http://auth-service:8000/users/auth"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	log.Println("Body: ", string(body))

	// rt.ID = primitive.NewObjectID()

	// err = nh.repo.CreateRegisteredStudentToTheCommpetition(rt)
	// if err != nil {
	// 	log.Println(err)
	// 	sendErrorWithMessage(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

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

func (nh *newHandler) GetDiplomaByUserId(w http.ResponseWriter, req *http.Request) {
	log.Println("Usli u GetDiplomaByUserId")

	// ctx := req.Context()
	// log.Println("Ctx: ", ctx)

	// authPayload, ok := ctx.Value("authorization_payload").(*token.Payload)
	// if !ok || authPayload == nil {
	// 	log.Println("AuthPayload: ", authPayload)
	// 	sendErrorWithMessage(w, "Authorization payload not found", http.StatusInternalServerError)
	// 	return
	// }

	// log.Println("Payload: ", authPayload)

	// id := authPayload.ID.Hex()
	// id = strings.Trim(id, "\"")
	// log.Println("Id: ", id)

	vars := mux.Vars(req)

	id := vars["id"]

	id = strings.Trim(id, "\"")
	log.Println("Id: ", id)

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
