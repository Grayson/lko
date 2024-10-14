package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("GET /", getRoot())
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

func getRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}
}
