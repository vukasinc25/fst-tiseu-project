package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vukasinc25/fst-tiseu-project/handler"
	"github.com/vukasinc25/fst-tiseu-project/middleware"
	"github.com/vukasinc25/fst-tiseu-project/repository"
	"github.com/vukasinc25/fst-tiseu-project/token"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	router := mux.NewRouter()
	router.StrictSlash(true)

	timeoutContext, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	tokenMaker, err := token.NewJWTMaker("12345678901234567890123456789012")
	if err != nil {
		log.Println("Ovde0: ", err)
	}

	newRepository, err := repository.New(timeoutContext)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer newRepository.Disconnect(timeoutContext)

	newRepository.Ping()

	server, err := handler.NewHandler(newRepository)
	if err != nil {
		log.Fatal(err)
		return
	}

	globalMiddleware := GlobalMiddleware(tokenMaker)
	router.Use(globalMiddleware)
	router.HandleFunc("/fakultet/create", server.CreateCompetition).Methods("POST")
	router.HandleFunc("/fakultet/user/create", server.CreateUser).Methods("POST")
	router.HandleFunc("/fakultet/user/registerToCompetition", server.CreateRegistrationUserToCompetition).Methods("POST")
	router.HandleFunc("/fakultet/user/diploma", server.CreateDiploma).Methods("POST")
	router.HandleFunc("/fakultet/user/diplomaByUserId", server.GetDiplomaByUserId).Methods("GET")
	router.HandleFunc("/fakultet/user/examResults", server.CreateUserExamResult).Methods("POST")
	router.HandleFunc("/fakultet/user/getResultsByCompetitionId/{id}", server.GetAllExamResultsByCompetitionId).Methods("GET")

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))
	srv := &http.Server{Addr: "0.0.0.0:8001", Handler: cors(router)}
	go func() {
		log.Println("server starting")
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-quit

	log.Println("service shutting down ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("server stopped")
}

func GlobalMiddleware(tokenMaker token.Maker) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return middleware.TokenMiddleware(tokenMaker)(next)
	}
}
