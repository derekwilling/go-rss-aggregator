package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code >= http.StatusInternalServerError {
		log.Println("Responding with 5XX error:", msg)
	}

	type errResp struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResp{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to marshal JSON response: %v", payload)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
