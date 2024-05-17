package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/vukasinc25/fst-tiseu-project/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
	AccessTokenKey          = "accessToken"
)

func TokenMiddleware(tokenMaker token.Maker) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract the token from the "Authorization" header
			tokenS := r.Header.Get("Authorization")
			if tokenS == "" {
				sendErrorWithMessage(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			tokenS = strings.Replace(tokenS, "Bearer ", "", 1)
			log.Println("Token: ", tokenS)

			// Send the token to the "/users/auth" path
			requestBody, err := json.Marshal(map[string]string{"token": tokenS})
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

			log.Println("Token: ", tokenS)

			payload, err := tokenMaker.VerifyToken(tokenS)
			if err != nil {
				// If the token verification fails, return an error
				writeError(w, http.StatusUnauthorized, err)
				return
			}

			log.Println("Payload: ", payload)

			// Store the payload in the request context
			ctx := context.WithValue(r.Context(), "authorization_payload", payload)
			r = r.WithContext(ctx)

			// Call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	}
}

func sendErrorWithMessage(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorResponse := map[string]string{"message": message}
	json.NewEncoder(w).Encode(errorResponse)
}

func writeError(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
}
