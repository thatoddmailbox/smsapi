package main

import (
	"encoding/json"
	"net/http"
)

type StatusResponse struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Error string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, thing interface{}) {
	w.Header().Set("Content-Type", "text/json")
	json.NewEncoder(w).Encode(thing)
}
