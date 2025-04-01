package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Sathish-30/authentication-and-authorization-golang/internal/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := models.BasicResponse{Message: "Welcome to Home"}
	json.NewEncoder(w).Encode(response)
}

func AdminsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := models.BasicResponse{Message: "List of admins"}
	json.NewEncoder(w).Encode(response)
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := models.BasicResponse{Message: "List of Users"}
	json.NewEncoder(w).Encode(response)
}
