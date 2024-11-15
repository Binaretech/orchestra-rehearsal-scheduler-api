package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
)

type Handler interface {
	Register(r *router.Router)
	RegisterProtected(group *router.Group)
}

func ResponseJson(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func ResponseError(w http.ResponseWriter, statusCode int, message string) {
	ResponseJson(w, statusCode, map[string]string{"error": message})
}

func ParseJsonBody(r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(data)
}
