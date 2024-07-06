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

	// tokenMaker, err := token.NewJWTMaker("12345678901234567890123456789012")
	// if err != nil {
	// 	log.Println("Ovde0: ", err)
	// }

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

	// globalMiddleware := GlobalMiddleware(tokenMaker)
	// router.Use(globalMiddleware)
	router.HandleFunc("/fakultet/createCompetition", server.CreateCompetition).Methods("POST")
	router.HandleFunc("/fakultet/competitions", server.GetAllCompetitions).Methods("GET")
	router.HandleFunc("/fakultet/competition/{id}", server.GetCompetitionById).Methods("GET")
	router.HandleFunc("/fakultet/user/create", server.CreateUser).Methods("POST")
	// router.HandleFunc("/fakultet/user/registerToCompetition", server.CreateRegistrationUserToCompetition).Methods("POST")
	router.HandleFunc("/fakultet/user/getRegistrationsToCompetition/{id}", server.GetAllRegistrationsToCompetition).Methods("GET")
	router.HandleFunc("/fakultet/user/registerToCompetition/{id}/{userId}", server.CreateRegistrationUserToCompetition).Methods("POST")
	router.HandleFunc("/fakultet/user/diplomaByUserId/{id}", server.GetDiplomaByUserId).Methods("GET") //ovde
	router.HandleFunc("/fakultet/user/examResults", server.CreateUserExamResult).Methods("POST")
	router.HandleFunc("/fakultet/user/getResultsByCompetitionId/{id}", server.GetAllExamResultsByCompetitionId).Methods("GET")
	router.HandleFunc("/fakultet/department", server.CreateDepartment).Methods("POST")
	router.HandleFunc("/fakultet/departments", server.GetAllDepartments).Methods("GET")
	router.HandleFunc("/fakultet/studyProgram", server.CreateStudyProgram).Methods("POST")
	router.HandleFunc("/fakultet/studyPrograms", server.GetAlltudyPrograms).Methods("GET")
	router.HandleFunc("/fakultet/studyProgram/{id}", server.GetStudyProgramById).Methods("GET")
	router.HandleFunc("/fakultet/diplomaRequest/{id}", server.DiplomaRequest).Methods("POST") //ovde
	router.HandleFunc("/fakultet/diplomaRequestsInPendingState", server.GetDiplomaRequestInPendingState).Methods("GET")
	router.HandleFunc("/fakultet/decideDiplomaReques/{id}", server.DecideDiplomaRequest).Methods("POST")
	router.HandleFunc("/fakultet/getDiplomaRequestsForUserId/{id}", server.GetDiplomaRequestsForUserId).Methods("GET") //ovde

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}),
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		gorillaHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}))
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
