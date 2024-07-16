package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondError(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	if err != nil {
		log.Printf("error serving %s: %s\n", r.URL.Path, err.Error())
	}

	w.WriteHeader(statusCode)
}

func respondJSON(w http.ResponseWriter, statusCode int, data any) {
	b, err := json.Marshal(data)
	if err != nil {
		log.Printf("error marshaling data: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
