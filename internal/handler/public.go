package handler

import "net/http"

func Public(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}
