package controllers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  bool
	Message string
	Data    interface{}
}

func respondJSON(w http.ResponseWriter, status int, msg string, payload interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	data := &Response{
		Status:  true,
		Message: msg,
		Data:    payload,
	}

	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, msg string) {

	w.WriteHeader(status)

	data := &Response{
		Status:  false,
		Message: msg,
		Data:    nil,
	}

	json.NewEncoder(w).Encode(data)
}
