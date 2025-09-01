package auth

import "github.com/gorilla/mux"

func CreateRoutes(r *mux.Router) {

	api := r.PathPrefix("/auth").Subrouter()

	api.HandleFunc("/create-nounce", createNounce).Methods("POST")
}
