package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespondWithJSON sends back a json response
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	log.SetPrefix("[jsonResponseError] :: ")
	response, err := json.Marshal(payload)

	if err != nil {
		log.Fatalf("Error => %v", err)
		RespondWithError(w, http.StatusInternalServerError, "Something went wrong on our server")
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		log.Fatalf("Error => %v", err)
	}
}

// RespondWithError sends back json response as an error
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"message": msg})
}
