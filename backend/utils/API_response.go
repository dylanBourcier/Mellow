package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	ErrorCode *string     `json:"errorCode"` // pointeur = null possible
}

// Réponse succès
func RespondJSON(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := APIResponse{
		Status:    "success",
		Message:   message,
		Data:      data,
		ErrorCode: nil,
	}

	json.NewEncoder(w).Encode(response)
}

// Réponse erreur
func RespondError(w http.ResponseWriter, code int, message string, errorCode string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := APIResponse{
		Status:    "error",
		Message:   message,
		Data:      nil,
		ErrorCode: &errorCode,
	}

	json.NewEncoder(w).Encode(response)
}
