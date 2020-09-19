package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func CreateServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", greetingHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
