package handler

import (
	"log"
	"net/http"

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

func (*UserHandler) Auth(w http.ResponseWriter, r *http.Request) {
	log.Println("Usli u Auth")
	w.Header().Add("Content-Type", "application/json")
	w.Write(r.TLS.OCSPResponse)
}
