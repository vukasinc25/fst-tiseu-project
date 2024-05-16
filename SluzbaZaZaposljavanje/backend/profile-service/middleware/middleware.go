package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func TokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the "Authorization" header
		token := r.Header.Get("Authorization")
		if token == "" {
			sendErrorWithMessage(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token = strings.Replace(token, "Bearer ", "", 1)
		log.Println("Token: ", token)

		// Send the token to the "/users/auth" path
		requestBody, err := json.Marshal(map[string]string{"token": token})
		if err != nil {
			sendErrorWithMessage(w, fmt.Sprintf("Error encoding request body: %v", err), http.StatusInternalServerError)
			return
		}

		// Send the token to the "/users/auth" path
		url := "http://auth-service:8000/users/auth"
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			sendErrorWithMessage(w, fmt.Sprintf("Error sending token: %v", err), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		nesto, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			sendErrorWithMessage(w, fmt.Sprintf("Error reading response body: %v", err), http.StatusInternalServerError)
			return
		}

		log.Println("Response: ", string(nesto))

		if strings.Contains(string(nesto), "token is invalid") {
			sendErrorWithMessage(w, "Token has expired", http.StatusUnauthorized)
			return
		}

		// Call the next handler in the chain
		next(w, r)
	}
}

func sendErrorWithMessage(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorResponse := map[string]string{"message": message}
	json.NewEncoder(w).Encode(errorResponse)
}
