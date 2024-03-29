package jsonutil

import (
	"encoding/json"
	"log"
	"net/http"
)

func DecodeRequestBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(v)
}

func RespondWithError(w http.ResponseWriter, err error, statusCode int, errMsg string) {
	log.Printf("Error: %s", err)
	if errMsg != "" {
		RespondWithJSON(w, statusCode, map[string]string{"error": errMsg})
	} else {
		w.WriteHeader(statusCode)
	}
}
