package main

import "net/http"

type Handler struct {
	db *Repo
}

func NewHandler(repo *Repo) (*Handler, error) {
	return &Handler{db: repo}, nil
}

func (hd *Handler) GetUserDiplomas(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(200)
}
