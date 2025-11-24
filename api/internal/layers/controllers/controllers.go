package controllers

import (
	"encoding/json"
	"net/http"
)

const (
	// success messages
	RESOURCE_CREATED_MSG string = "resource created successfully"

	// error messages
	INVALID_REQUEST_MSG      string = "invalid request"
	SOMETHING_WENT_WRONG_MSG string = "something went wrong"
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
