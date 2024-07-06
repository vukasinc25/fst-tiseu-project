package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	router := mux.NewRouter()
	router.StrictSlash(true)

	//timeoutContext, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	//defer cancel()

	repository, err := NewRepository()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer repository.Disconnect()

	err = repository.CreateData()
	if err != nil {
		log.Fatal(err)
		return
	}

	server, err := NewHandler(repository)
	if err != nil {
		log.Fatal(err)
		return
	}
	router.HandleFunc("/skola/diplomas", server.GetUserDiplomas).Methods("POST")
	// router.Use(GlobalMiddleware)
	//router.HandleFunc("/fakultet/create", server.CreateCompetition).Methods("POST")

	srv := &http.Server{Addr: "0.0.0.0:8005", Handler: router}
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
