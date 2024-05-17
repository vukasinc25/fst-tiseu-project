package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vukasinc25/fst-tiseu-project/handler"
	"github.com/vukasinc25/fst-tiseu-project/repository"
	"github.com/vukasinc25/fst-tiseu-project/token"
)

func main() {
	port := "8000"
	timeoutContext, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	// Initialize Gorilla Mux router and CORS middleware
	router := mux.NewRouter()
	router.StrictSlash(true)
	//cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	// Initialize loggers with prefixes for different components
	// logger := log.New(os.Stdout, "[auth-api] ", log.LstdFlags)
	// storeLogger := log.New(os.Stdout, "[auth-store] ", log.LstdFlags)

	// Create a JWT token maker
	tokenMaker, err := token.NewJWTMaker("12345678901234567890123456789012")
	if err != nil {
		log.Println("Ovde0: ", err)
	}

	log.Println("Token: ", tokenMaker)

	// NoSQL: Initialize auth Repository store
	store, err := repository.New(timeoutContext, "", "", "")
	if err != nil {
		log.Println("Ovde1: ", err)
	}
	defer store.Disconnect(timeoutContext)

	// // Check if the data store connection was established
	store.Ping()

	// Create a user handler service
	service := handler.NewUserHandler(store, tokenMaker)
	if err != nil {
		log.Println("Ovde2: ", err)
	}

	log.Println("Ovde3: ", service)
	router.HandleFunc("/users/auth", service.Auth).Methods("POST")
	router.HandleFunc("/users/create", service.CreateUser).Methods("POST")
	router.HandleFunc("/users/login", service.LoginUser).Methods("POST	")

	// Enable CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow all origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	}).Handler(router)

	log.Println("Cors: ", corsHandler)

	// Configure the HTTP server
	server := &http.Server{Addr: ":8000", Handler: corsHandler}

	// Print a message indicating the server is listening
	log.Println("Server listening on port", port)

	// Start the HTTP server in a goroutine
	go func() {
		log.Println("server starting")
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	// Listen for signals to gracefully shut down the server
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT)
	signal.Notify(sigCh, syscall.SIGKILL)

	sig := <-sigCh
	log.Println("Received terminate, graceful shutdown", sig)

	// Create a new context for graceful shutdown with a timeout of 30 seconds
	timeoutContext, _ = context.WithTimeout(context.Background(), 30*time.Second)

	// Attempt to gracefully shut down the server
	if server.Shutdown(timeoutContext) != nil {
		log.Println("Cannot gracefully shutdown...")
	}
	log.Println("Server stopped")
}

func loadConfig() map[string]string {
	config := make(map[string]string)
	config["conn_service_address"] = fmt.Sprintf("http://%s:%s", os.Getenv("PROF_SERVICE_HOST"), os.Getenv("PROF_SERVICE_PORT"))
	return config
}

func CreateUser(w http.ResponseWriter, req *http.Request) {
	sendErrorWithMessage(w, "User Created", http.StatusCreated)
}

func sendErrorWithMessage(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorResponse := map[string]string{"message": message}
	json.NewEncoder(w).Encode(errorResponse)
}
