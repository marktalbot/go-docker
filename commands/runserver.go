package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Welcome to this life-changing API. Boop!")
	})

	fmt.Println("Server listening!")
	http.ListenAndServe(":80", r)
}
