package api

import (
    "net/http"
    "encoding/json"
)

// Response represents the structure of the API response
type Response struct {
    Message string `json:"message"`
}

// GetHandler handles GET requests
func GetHandler(w http.ResponseWriter, r *http.Request) {
    response := Response{Message: "GET request successful"}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

// PostHandler handles POST requests
func PostHandler(w http.ResponseWriter, r *http.Request) {
    response := Response{Message: "POST request successful"}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}