package main

import (
	"context"
	"github.com/gorilla/mux"
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

	newRepository, err := repository.New(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	server, err := handler.NewHandler(newRepository)
	if err != nil {
		log.Fatal(err)
		return
	}

	router.HandleFunc("/", server.CreateUser).Methods("POST")

	srv := &http.Server{Addr: "0.0.0.0:8003", Handler: router}
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
