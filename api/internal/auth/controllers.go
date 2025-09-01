package auth

import "net/http"

func createNounce(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello!"))
}
