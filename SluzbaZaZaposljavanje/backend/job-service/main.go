package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/vukasinc25/fst-tiseu-project/handler"
	"github.com/vukasinc25/fst-tiseu-project/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	router := mux.NewRouter()
	router.StrictSlash(true)

	logger := log.New(os.Stdout, "[jobListing-api] ", log.LstdFlags)

	newRepository, err := repository.New(context.Background(), logger)
	if err != nil {
		log.Fatal(err)
		return
	}

	server, err := handler.NewHandler(logger, newRepository)
	if err != nil {
		log.Fatal(err)
		return
	}

	router.HandleFunc("/createJobListing", server.CreateJobListing).Methods("POST")
	router.HandleFunc("/applyForJob", server.CreateJobApplication).Methods("POST")
	router.HandleFunc("/getJobListings", server.GetAllJobListings).Methods("GET")

	// Enable CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow all origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	}).Handler(router)

	srv := &http.Server{Addr: "0.0.0.0:8012", Handler: corsHandler}
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

	// gracefully stop server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("server stopped")
}
