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

	"github.com/vukasinc25/fst-tiseu-project/model"
	"github.com/vukasinc25/fst-tiseu-project/repository"
	"github.com/vukasinc25/fst-tiseu-project/token"
)

type UserHandler struct {
	db       *repository.UserRepo
	jwtMaker token.Maker
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(r *repository.UserRepo, jwtMaker token.Maker) *UserHandler {
	return &UserHandler{r, jwtMaker}
}

func writeError(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
}

func (uh *UserHandler) Auth(w http.ResponseWriter, r *http.Request) {
	log.Println("req received")

	dec := json.NewDecoder(r.Body)

	var rt model.ReqToken
	err := dec.Decode(&rt)
	if err != nil {
		log.Println(err)
		log.Println("Request decode error")
	}

	log.Println(rt.Token)

	payload, err := uh.jwtMaker.VerifyToken(rt.Token)
	if err != nil {
		// If the token verification fails, return an error
		log.Println("error in token verification")
		writeError(w, http.StatusUnauthorized, err)
		return
	}

	respBytes, err := json.Marshal(payload.ID)
	if err != nil {
		log.Println("error while creating response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(respBytes)

}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, req *http.Request) {
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

	err = uh.db.Insert(rt, ctx)
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
