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
	"github.com/vukasinc25/fst-tiseu-project/handler"
	"github.com/vukasinc25/fst-tiseu-project/repository"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	router := mux.NewRouter()
	router.StrictSlash(true)

	timeoutContext, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

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


	router.HandleFunc("aprcroso/createFirm", server.CreateFirm).Methods("POST")

	srv := &http.Server{Addr: "0.0.0.0:8004", Handler: router}
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

