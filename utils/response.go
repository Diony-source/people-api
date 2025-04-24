package utils

import (
	"encoding/json"
	"net/http"
)

type successResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// JSON returns a standard success JSON response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(successResponse{
		Success: true,
		Data:    data,
	})
}

// Error returns a standard error JSON response
func Error(w http.ResponseWriter, statusCode int, msg string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errText := ""
	if err != nil {
		errText = err.Error()
	}

	json.NewEncoder(w).Encode(errorResponse{
		Success: false,
		Message: msg,
		Error:   errText,
	})
}
