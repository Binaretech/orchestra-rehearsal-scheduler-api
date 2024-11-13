package api

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func ResponseError(w http.ResponseWriter, statusCode int, message string) {
	ResponseJson(w, statusCode, map[string]string{"error": message})
}
